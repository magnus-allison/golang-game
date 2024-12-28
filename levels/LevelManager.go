package levels

type LevelManager struct{
	Levels []*Level
	Level *Level
}

func CreateLevelManager() *LevelManager {
	levelOne := createLevelOne()
	return &LevelManager{
		Levels: []*Level{
			levelOne,
		},
		Level: levelOne,
	}
}

func (lm *LevelManager) LoadLevel(level *Level) {
	lm.Level = level
}
