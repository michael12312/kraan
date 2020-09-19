//Package tarconsumer provides an interface for processing repositories.
//go:generate mockgen -destination=../mocks/tarconsumer/mockTarconsumer.go -package=mocks -source=tarconsumer.go . TarConsumer
package tarconsumer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/fluxcd/pkg/untar"
)

type TarConsumer interface {
	SetCtx(ctx context.Context)
	SetURL(utl string)
	SetHTTPClient(httpClient *http.Client)
	GetTar(ctx context.Context) ([]byte, error)
}

type tarConsumerData struct {
	ctx         context.Context
	url         string
	httpClient  *http.Client
	TarConsumer `json:"-"`
}

func NewTarConsumer(ctx context.Context, client *http.Client, url string) TarConsumer {
	return &tarConsumerData{
		ctx:        ctx,
		url:        url,
		httpClient: client,
	}
}

func (t *tarConsumerData) SetCtx(ctx context.Context) {
	t.ctx = ctx
}

func (t *tarConsumerData) SetHTTPClient(httpClient *http.Client) {
	t.httpClient = httpClient
}

func (t *tarConsumerData) SetURL(url string) {
	t.url = url
}

func (t *tarConsumerData) GetTar(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, t.url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return streamToByte(resp.Body)
}

func UnpackTar(tar []byte, path string) (err error) {
	r := bytes.NewReader(tar)
	_, err = untar.Untar(r, path)
	return err
}

func streamToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(stream)
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, fmt.Errorf("tar data is empty")
	}
	return buf.Bytes(), nil
}
