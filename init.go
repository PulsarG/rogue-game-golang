package main

func initNewGame(hp, isNewBackPack int) *Session {

	ns := Session{}
	ns.GenMap()
	ns.Player = PlayerInit(&ns.Rooms[ns.StartRoomIdx], hp)
	ns.CurrentLvl = GameLvl
	ns.EnemysListInit()
	//s.CurrentState = StateMainMenu
	ns.EndRoom = ns.EndRoomInit()
	ns.GoldInit()
	ns.Statistic = statisticInit()
	if isNewBackPack == 0 {
		ns.Backpack = InitBackpack()
	} else {
		ns.Backpack = s.Backpack
		ns.Player.Weapon = s.Player.Weapon
	}
	ns.FlasksInit()
	ns.ScrollsInit()
	ns.WeaponsInit()
	ns.FoodInit()
	ns.PlayerScanFov(Visible)
	return &ns
}

func statisticInit() *Statistic {
	return &Statistic{
		RunNum: s.readNumRun(),
	}
}
