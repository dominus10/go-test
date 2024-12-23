package registry

type SceneNavigator interface {
	SceneNext(filename string)
	SceneReturn()
}

type InventoryManipulator interface {
	ItemAdd(id string, count int)
	ItemUse(id string, count int)
	ItemUseOn(target string, id string, count int)
	ItemDrop(id string, count int)
	ItemSteal(id string, count int)
}

type StatusManipulator interface {
	StatusRaise(target string, stat string, count int)
	StatusDrop(target string, stat string, count int)
	StatusRaiseTemp(target string, stat string, count int)
	StatusDropTemp(target string, stat string, count int)
}

type BattleActions interface {
	BattleAttack(target string, types string)
	BattleAction(target string, types string)
	BattleSpell(target string, types string)
	BattleFlee()
	BattleSurrender()
}