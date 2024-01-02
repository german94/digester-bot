package gptassistant

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type Assistant interface {
	Init(ctx context.Context, fileContent []byte, fileName string) error
	Ask(ctx context.Context, question string) (string, error)
}

type GPTAssistant struct {
	client      *openai.Client
	assistantID string
	threadID    string
	key         string
}

func New(key string) *GPTAssistant {
	return &GPTAssistant{
		client: openai.NewClient(key),
		key:    key,
	}
}

func (gpta *GPTAssistant) Init(ctx context.Context, fileContent []byte, fileName string) error {
	// f, err := gpta.client.CreateFile(ctx, openai.FileRequest{
	// 	FilePath: filepath,
	// 	Purpose:  string(openai.PurposeAssistants),
	// })
	f, err := gpta.uploadFromMemory(fileContent, fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	log.Println("fileid: " + f)

	resp, err := gpta.client.CreateAssistant(ctx, openai.AssistantRequest{
		Model: openai.GPT4TurboPreview,
		Instructions: func() *string {
			i := fmt.Sprintf("You are a tool to help people study the subject in the file attached (%v). You might be asked questions or tasks like summarization or explanations.",
				f)
			return &i
		}(),
		Tools:   []openai.AssistantTool{{Type: openai.AssistantToolTypeRetrieval}},
		FileIDs: []string{f},
	})
	if err != nil {
		return fmt.Errorf("error creating assistant: %w", err)
	}
	gpta.assistantID = resp.ID
	log.Println("assistantid: " + resp.ID)

	t, err := gpta.client.CreateThread(ctx, openai.ThreadRequest{})
	if err != nil {
		return fmt.Errorf("error creating thread: %w", err)
	}
	gpta.threadID = t.ID

	return nil
}

func (gpta *GPTAssistant) uploadFromMemory(fileData []byte, fileName string) (string, error) {
	url := "https://api.openai.com/v1/files"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if err := writer.WriteField("purpose", "assistants"); err != nil {
		return "", err
	}
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, bytes.NewReader(fileData))
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+gpta.key)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error uploading file: http %v status code", resp.StatusCode)
	}

	var data map[string]any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	id, ok := data["id"].(string)
	if !ok {
		return "", fmt.Errorf("error mapping response: %+v", data)
	}

	return id, nil
}

func (gpta *GPTAssistant) Ask(ctx context.Context, question string) (string, error) {
	log.Printf("Asking question: %v\n", question)
	if gpta.assistantID == "" || gpta.threadID == "" {
		return "", fmt.Errorf("assistant not initialized yet")
	}
	if _, err := gpta.client.CreateMessage(ctx, gpta.threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	}); err != nil {
		return "", fmt.Errorf("error creating message: %w", err)
	}
	run, err := gpta.client.CreateRun(ctx, gpta.threadID, openai.RunRequest{
		AssistantID: gpta.assistantID,
	})
	if err != nil {
		return "", fmt.Errorf("error creating run: %w", err)
	}
	for {
		time.Sleep(5 * time.Second)
		run, err := gpta.client.RetrieveRun(ctx, gpta.threadID, run.ID)
		if err != nil {
			return "", fmt.Errorf("error retrieving run: %w", err)
		}
		if run.Status == openai.RunStatusCompleted {
			break
		}
	}
	msgs, err := gpta.client.ListMessage(ctx, gpta.threadID, nil, nil, nil, nil)
	if err != nil {
		return "", fmt.Errorf("error retrieving messages: %w", err)
	}
	if len(msgs.Messages) == 0 {
		return "", fmt.Errorf("no messages: %+v", msgs)
	}
	contentArr := msgs.Messages[0].Content
	if len(contentArr) == 0 {
		return "", fmt.Errorf("no content in msg: %+v", msgs)
	}
	if contentArr[0].Text == nil {
		return "", fmt.Errorf("no text in message: %+v", msgs)
	}

	return contentArr[0].Text.Value, nil
}
