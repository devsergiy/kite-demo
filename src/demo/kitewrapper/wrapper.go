package kitewrapper

import (
	conf "demo/config"
	"fmt"
	"net/url"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/protocol"
)

// Helper class with methods to register and dial other kites
type Wrapper struct {
	*kite.Kite
	cfg *conf.KiteConfig
}

func NewKiteWrapper(cfg *conf.KiteConfig) *Wrapper {
	k := kite.New(cfg.Name, cfg.Version)
	k.Config = config.MustGet()
	k.Config.Port = cfg.Port

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
		Name: name,
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
