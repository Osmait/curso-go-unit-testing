package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePokemonSuccess(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("sample/pokeapi_reponse.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)
	ParsedPokemon, err := ParsePokemon(response)
	c.NoError(err)

	body, err = ioutil.ReadFile("sample/api_response.json")
	c.NoError(err)

	var expectedPokemon models.Pokemon

	err = json.Unmarshal([]byte(body), &expectedPokemon)
	c.NoError(err)

	c.Equal(expectedPokemon, ParsedPokemon)

}
func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)
	body, err := ioutil.ReadFile("sample/pokeapi_reponse.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	response.PokemonType = []models.PokemonType{}
	_, err = ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(ErrNotFoundPokemonType, err.Error())
}
