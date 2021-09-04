package tenniskata

type tennisGame1 struct {
	m_score1, m_score2  int
	player1Name, player2Name string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1 {
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.m_score1 == game.m_score2 {
		switch game.m_score1 {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if game.m_score1 >= 4 || game.m_score2 >= 4 {
		switch minusResult := (game.m_score1 - game.m_score2); {
		case minusResult == 1:
			score = "Advantage player1"
		case minusResult == -1:
			score = "Advantage player2"
		case minusResult >= 2:
			score = "Win for player1"
		default:
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.m_score1
			} else {
				score += "-"
				tempScore = game.m_score2
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}
