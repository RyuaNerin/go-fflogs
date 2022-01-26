package fflogs

import (
	"context"

	"fflogs/structure"
)

// Gets an array of CharacterRanking objects. Each CharacterRanking corresponds to a single rank on a fight for the specified character.
func (c *Client) RankingCharacter(context context.Context, opt *RankingCharacterOptions) ([]structure.CharacterRanking, error) {
	var resp []structure.CharacterRanking
	err := c.Raw.RankingCharacter(context, opt, &resp)
	return resp, err
}
