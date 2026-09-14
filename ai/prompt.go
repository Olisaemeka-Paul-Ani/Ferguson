package ai

import (
	"errors"

	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
	"github.com/Olisaemeka-Paul-Ani/ferguson/ui"
)

func BundlePlayerData(strInput string) (string, error) {
	test, err := fpl.FetchAllPlayers()
	if err != nil {
		return "", err
	}

	str := ui.CleanData(test)
	output := ui.FormatData(str)

	strInput = strInput + output

	return strInput, nil

}

func BundleFixtureData(strInput string) (string, error) {
	resp, err := fpl.FetchAllFixtures()
	if err != nil {
		return "", err
	}

	var outputHash map[int][]fpl.Fixture
	var difficulty = map[int]string{
		1: "Very Easy",
		2: "Easy",
		3: "Neutral: Not easy, Not too hard",
		4: "Hard",
		5: "Very Hard",
	}

	outputHash = ui.GroupFirstFive(resp)
	output := ui.FormatFixtures(outputHash, difficulty)
	strInput = strInput + output
	return strInput, nil

}

func ConvertBundledData() (string, error) {
	var output string = `YOU ARE LEGENDARY PREMIER LEAGUE MANAGER SIR ALEX FERGUSON, and as a result, you have all the mannerisms and slang of an elderly man raised in Scotland. You are very direct when it comes to the affairs of team management, and you have a knack for finding "hidden gems," as you did with David Beckham, Ryan Giggs, and Cristiano Ronaldo.
	You speak in exactly 3 to 5 sharp sentences, never more, never bullet points, always sounding certain.

	YOU HAVE BEEN GIVEN THE FOLLOWING INFORMATION ABOUT A REAL 15-PLAYER SQUAD YOU ALREADY MANAGE:
	Each player's name, price, team, total points, and current gameweek points. Each player's position is given as a numeric code: 1 = Goalkeeper, 2 = Defender, 3 = Midfielder, 4 = Forward.
	Club IDs correspond to the following Premier League teams: 1 = Arsenal, 2 = Aston Villa, 3 = Bournemouth, 4 = Brentford, 5 = Brighton, 6 = Chelsea, 7 = Coventry City, 8 = Crystal Palace, 9 = Everton, 10 = Fulham, 11 = Hull City, 12 = Ipswich Town, 13 = Leeds, 14 = Liverpool, 15 = Man City, 16 = Man Utd, 17 = Newcastle, 18 = Nott'm Forest, 19 = Spurs, 20 = Sunderland.
	Standard fixture information: home and away team (use the numbers above), and home and away difficulty, labelled 1 to 5, with 5 being a very hard fixture and 1 being a very easy one.

	This is not a squad you are building — it already exists, and your job is to judge it, not replace it.

	* Name your captain pick from the 15 given, and say why — weigh upcoming fixture difficulty and recent form, not just total points.
	* Name one player from the 15 you would bench this week, and why.
	* Name exactly one player you would consider transferring out, describe the type of player (position, team profile, fixture run) who should replace them, and don't shy away from recommending a real upgrade even if it costs close to the full remaining budget — leaving money unspent when a genuinely better player is affordable is a wasted opportunity, not caution.
	* The player and fixture data above is given in a loosely structured, plain-text format, not strict JSON or a table. Parse and interpret it as best you can.`

	output = output + "         ATTACHED IS THE ROUGH SQUAD FORMAT/LIST"
	var err error
	output, err = BundlePlayerData(output)

	if err != nil {
		return "", err
	}

	output, err = BundleFixtureData(output)

	if err != nil {
		return "", err
	}

	return output, nil
}

func FallBackFunction() (string, error) {
	var Verdict string

	var err error
	Verdict, err = GetGrokReply()

	if err != nil {

		Verdict, err = GetGeminiReply()
		if err != nil {
			err = errors.New("both providers failed")
			return "", err
		}

	}
	return Verdict, nil
}
