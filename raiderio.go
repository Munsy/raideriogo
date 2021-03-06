package raideriogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// RaiderIOResponse is an empty interface.
type RaiderIOResponse interface {
}

// RaiderIOClient retrieves data from the Raider.io API.
type RaiderIOClient struct {
	client *http.Client
}

// New creates a new RaiderIOClient.
func New() *RaiderIOClient {
	r := &RaiderIOClient{
		&http.Client{
			Timeout: time.Second * 10,
		},
	}

	return r
}

// GetCharacterProfile creates a new ViewCharacterProfileResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetCharacterProfile(region, realm, name, fields string) (*ViewCharacterProfileResponse, error) {
	var characterProfile = &ViewCharacterProfileResponse{}

	err := rio.get(EndpointCharacter(region, realm, name, fields), characterProfile)

	if nil != err {
		return nil, err
	}

	return characterProfile, nil
}

// GetGuildProfile creates a new ViewGuildProfileResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetGuildProfile(region, realm, name, fields string) (*ViewGuildProfileResponse, error) {
	var guildProfile = &ViewGuildProfileResponse{}

	err := rio.get(EndpointGuild(region, realm, name, fields), guildProfile)

	if nil != err {
		return nil, err
	}

	return guildProfile, nil
}

// GetMythicPlusAffixes creates a new ViewMythicPlusAffixesResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetMythicPlusAffixes(region, locale string) (*ViewMythicPlusAffixesResponse, error) {
	var affixes = &ViewMythicPlusAffixesResponse{}

	err := rio.get(EndpointMythicPlusAffixes(region, locale), affixes)

	if nil != err {
		return nil, err
	}

	return affixes, nil
}

// GetMythicPlusRuns creates a new ViewMythicPlusRunsResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetMythicPlusRuns(season, region, dungeon, affixes string) (*ViewMythicPlusRunsResponse, error) {
	var runs = &ViewMythicPlusRunsResponse{}

	err := rio.get(EndpointMythicPlusRuns(season, region, dungeon, affixes), runs)

	if nil != err {
		return nil, err
	}

	return runs, nil
}

// GetRaidingBossRankings creates a new ViewRaidingBossRankingsResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetRaidingBossRankings(raid, boss, difficulty, region, realm string) (*ViewRaidingBossRankingsResponse, error) {
	var bossRankings = &ViewRaidingBossRankingsResponse{}

	err := rio.get(EndpointRaidingBossRankings(raid, boss, difficulty, region, realm), bossRankings)

	if nil != err {
		return nil, err
	}

	return bossRankings, nil
}

// GetRaidingHallOfFame creates a new ViewRaidingHallOfFameResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetRaidingHallOfFame(raid, difficulty, region string) (*ViewRaidingHallOfFameResponse, error) {
	var hallOfFame = &ViewRaidingHallOfFameResponse{}

	err := rio.get(EndpointRaidingHallOfFame(raid, difficulty, region), hallOfFame)

	if nil != err {
		return nil, err
	}

	return hallOfFame, nil
}

// GetRaidingProgression creates a new ViewRaidingProgressionResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetRaidingProgression(raid, difficulty, region string) (*ViewRaidingProgressionResponse, error) {
	var progression = &ViewRaidingProgressionResponse{}

	err := rio.get(EndpointRaidingProgression(raid, difficulty, region), progression)

	if nil != err {
		return nil, err
	}

	return progression, nil
}

// GetRaidingRaidRankings creates a new ViewRaidingRaidRankingsResponse by calling the Raider.io API.
func (rio *RaiderIOClient) GetRaidingRaidRankings(raid, difficulty, region, realm string) (*ViewRaidingRaidRankingsResponse, error) {
	var rankings = &ViewRaidingRaidRankingsResponse{}

	err := rio.get(EndpointRaidingRaidRankings(raid, difficulty, region, realm), rankings)

	if nil != err {
		return nil, err
	}

	return rankings, nil
}

// Returns an interface corresponding to the given endpoint.
func (rio *RaiderIOClient) get(endpoint string, raiderResponse RaiderIOResponse) error {
	httpResponse, err := rio.client.Get(endpoint)

	if nil != err {
		return err
	}
	defer httpResponse.Body.Close()

	data, err := ioutil.ReadAll(httpResponse.Body)
	if nil != err {
		return err
	}

	return json.Unmarshal(data, &raiderResponse)
}
