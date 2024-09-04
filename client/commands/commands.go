package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/mukappalambda/my-trader/client/common"
	pb "github.com/mukappalambda/my-trader/internal/adapters/grpc/message/v1"
	"github.com/mukappalambda/my-trader/internal/adapters/rest/types"
	"github.com/spf13/cobra"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunApply(cmd *cobra.Command, args []string) {
	filename, _ := cmd.Flags().GetString("filename")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %q: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()
	byt, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var schema types.Schema
	err = json.Unmarshal(byt, &schema)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", schema)

	srUrl, _ := cmd.Flags().GetString("schema-registry-url")
	buf, err := json.Marshal(schema)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/schemas", srUrl), bytes.NewBuffer(buf))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}

func RunGenerate(cmd *cobra.Command, args []string) {
	byt, err := json.MarshalIndent(types.DefaultSchema, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to serialize schema: %v\n", err)
	}
	output, _ := cmd.Flags().GetString("output")
	if output == "" {
		fmt.Printf("%s\n", byt)
		return
	}
	file, err := os.Create(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open the file: %v\n", err)
	}
	defer file.Close()
	_, err = file.Write(byt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write out the file: %v\n", err)
	}
}

func RunSend(cmd *cobra.Command, args []string) {
	topic, _ := cmd.Flags().GetString("topic")
	message, _ := cmd.Flags().GetString("message")
	serverUrl, _ := cmd.Flags().GetString("server-url")
	srUrl, _ := cmd.Flags().GetString("schema-registry-url")
	schemaName, _ := cmd.Flags().GetString("schema")

	params := url.Values{}
	params.Set("name", schemaName)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/schemas?%s", srUrl, params.Encode()), nil)
	common.PrintToStderrThenExit(err)
	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}
	resp, err := client.Do(req)
	common.PrintToStderrThenExit(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	common.PrintToStderrThenExit(err)
	var schema types.Schema
	err = json.Unmarshal(body, &schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not deserialize to the given schema: %v\n", err)
		os.Exit(1)
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(message), &result)
	common.PrintToStderrThenExit(err)
	collected := make(map[string]interface{})
	for _, field := range schema.Fields {
		if field.Type == "double" {
			if value, ok := result[field.Name].(float64); ok {
				collected[field.Name] = value
			}
		}
		if field.Type == "string" {
			if value, ok := result[field.Name].(string); ok {
				collected[field.Name] = value
			}
		}
	}
	conn, err := grpc.NewClient(serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.PutMessage(ctx, &pb.MessageRequest{
		Topic:     topic,
		Message:   message,
		CreatedAt: toDateTime(time.Now()),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Sent message: schema-subject=%s schema-name=%s ", schema.Subject, schema.Name)
	for k, v := range collected {
		switch typedV := v.(type) {
		case float64:
			fmt.Printf("%s=%f ", k, typedV)
		case string:
			fmt.Printf("%s=%s ", k, typedV)
		}
	}
	fmt.Printf("\n")
}

func RunGet(cmd *cobra.Command, args []string) {
	srUrl, _ := cmd.Flags().GetString("schema-registry-url")
	subject, _ := cmd.Flags().GetString("subject")
	name, _ := cmd.Flags().GetString("name")

	k := "subject"
	v := subject
	if name != "" {
		k = "name"
		v = name
	}

	params := url.Values{}
	params.Set(k, v)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/schemas?%s", srUrl, params.Encode()), nil)
	common.PrintToStderrThenExit(err)
	client := &http.Client{
		Timeout: 200 * time.Millisecond,
	}
	resp, err := client.Do(req)
	common.PrintToStderrThenExit(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	common.PrintToStderrThenExit(err)
	fmt.Println(string(body))
}

func RunCheck(cmd *cobra.Command, args []string) {
	filename, _ := cmd.Flags().GetString("filename")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open %q: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()
	byt, err := io.ReadAll(file)
	common.PrintToStderrThenExit(err)

	var schema types.Schema
	err = json.Unmarshal(byt, &schema)
	common.PrintToStderrThenExit(err)
	if schema.Name == "" || schema.Subject == "" {
		fmt.Fprintln(os.Stderr, `Field names "schema" and "subject" cannot be empty.`)
		os.Exit(1)
	} else {
		fmt.Println("Checked!")
	}
}

func toDateTime(t time.Time) *datetime.DateTime {
	return &datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}
