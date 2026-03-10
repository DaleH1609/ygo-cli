package connections

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DaleH1609/ygo-cli/structs"
)

func FetchCard(name string) (structs.Card, error) {
	url := fmt.Sprintf("https://db.ygoprodeck.com/api/v7/cardinfo.php?name=%s", name)

	resp, err := http.Get(url)
	if err != nil {
		return structs.Card{}, fmt.Errorf("failed to reach API: %w", err)
	}
	defer resp.Body.Close()

	var cardResponse structs.CardResponse
	err = json.NewDecoder(resp.Body).Decode(&cardResponse)
	if err != nil {
		return structs.Card{}, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(cardResponse.Data) == 0 {
		return structs.Card{}, fmt.Errorf("no card found with name: %s", name)
	}

	return cardResponse.Data[0], nil
}
