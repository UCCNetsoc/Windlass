package config

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/UCCNetworkingSociety/Windlass/app/connections"
	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

func setSharedSecret() error {
	client, err := connections.GetConsul()
	if err != nil {
		return err
	}

	kv, _, err := client.KV().Get(viper.GetString("consul.path")+"/secret", &api.QueryOptions{
		RequireConsistent: true,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("error getting shared secret: %v", err))
	}

	if kv != nil {
		log.Debug("found shared secret in Consul")
		viper.Set("windlass.secret", string(kv.Value))
		return nil
	}

	log.Debug("generating and setting shared secret in Consul")

	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		return errors.New(fmt.Sprintf("error generating shared secret: %v", err))
	}

	encodedKey := hex.EncodeToString(key)

	if _, err := client.KV().Put(&api.KVPair{
		Key:   viper.GetString("consul.path") + "/secret",
		Value: []byte(encodedKey),
	}, &api.WriteOptions{}); err != nil {
		return errors.New(fmt.Sprintf("error setting shared secret: %v", err))
	}

	viper.Set("windlass.secret", string(encodedKey))

	return nil
}
