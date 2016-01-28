package kitewrapper

import (
	"demo/config"
	"fmt"
	"net/url"

	"github.com/koding/kite"
	"github.com/koding/kite/kitekey"
	"github.com/koding/kite/protocol"
)

// Helper class with methods to register and dial other kites
type Wrapper struct {
	*kite.Kite
	cfg config.KiteConfig
}

func NewKiteWrapper(KontrolURL, KontrolUser string, cfg config.KiteConfig) *Wrapper {
	k := kite.New(cfg.Name, cfg.Version)
	k.Config.Port = cfg.Port
	k.Config.KiteKey, _ = kitekey.Read()
	k.Config.KontrolURL = KontrolURL
	k.Config.KontrolUser = KontrolUser

	return &Wrapper{k, cfg}
}

// Register current kite in Kontrol
func (w *Wrapper) RegisterToKontrol() error {
	w.GetKey()

	err := w.RegisterForever(&url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", w.cfg.Host, w.cfg.Port),
		Path:   w.cfg.Path,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Finds kite by name and connects to it
func (w *Wrapper) FindAndDial(name string) (*kite.Client, error) {
	kites, err := w.GetKites(&protocol.KontrolQuery{
		Name: w.cfg.Name,
	})
	if err != nil {
		return nil, err
	}

	k := kites[0]
	err = k.Dial()

	if err != nil {
		return nil, err
	}

	return k, nil
}
