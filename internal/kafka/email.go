package kafka

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type mqEmail struct {
	Address string `json:"address"`
	Content string `json:"content"`
}

func MakeAcceptanceEmail(email string, baseUrl string, taskId int, letterUuid uuid.UUID) ([]byte, error) {
	content := []string{
		fmt.Sprintf("TaskID: %v", taskId),
		fmt.Sprintf("To accept task %v/tasks/accept/%v", baseUrl, letterUuid),
		fmt.Sprintf("To decline task %v/tasks/decline/%v", baseUrl, letterUuid),
	}

	msg := mqEmail{
		Address: email,
		Content: strings.Join(content, " <br>"),
	}

	rawMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return rawMsg, nil
}

func MakeCancellationEmail(email string, taskId int) ([]byte, error) {
	content := []string{
		fmt.Sprintf("TaskID: %v", taskId),
		"Task declined.",
	}

	msg := mqEmail{
		Address: email,
		Content: strings.Join(content, " <br>"),
	}

	rawMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return rawMsg, nil
}
