/*
 * Copyright (c) 2024 Ruiyuan "mizumoto-cn" Xu
 *
 * This file is part of "github.com/mizumoto-cn/mizubot".
 *
 * Licensed under the Mizumoto General Public License v1.5 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://github.com/mizumoto-cn/mizubot/blob/main/LICENSE
 *     https://github.com/mizumoto-cn/mizubot/blob/main/licensing
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package post

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type IPoster interface {
	Post(ctx context.Context) error
}

type poster struct {
	WebhookURL string
	Content    string
	Channel1   string
	Channel2   string
	User1Mail  string
	User2Mail  string
}

// Post sends a message to the specified URL
func (p *poster) Post(ctx context.Context) error {
	// Create the message payload
	payload := map[string]string{
		"at_user_1":    p.User1Mail,
		"to_channel_1": p.Channel1,
		"to_mizumoto":  p.Content,
		"at_user_2":    p.User2Mail,
		"content":      p.Content,
		"to_channel_2": p.Channel2,
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create a new POST request
	req, err := http.NewRequestWithContext(ctx, "POST", p.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to post message, request: %s, status code: %d\n", string(jsonData), resp.StatusCode)
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(b)
		return fmt.Errorf("failed to post message, status code: %d", resp.StatusCode)
	}

	return nil
}

func NewPoster(webhookURL, content, channel1, channel2, user1Mail, user2Mail string) IPoster {
	return &poster{
		WebhookURL: webhookURL,
		Content:    content,
		Channel1:   channel1,
		Channel2:   channel2,
		User1Mail:  user1Mail,
		User2Mail:  user2Mail,
	}
}
