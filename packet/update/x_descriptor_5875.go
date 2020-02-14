package update

func init5875() {
	dc := NewDescriptorCompiler(5875)

	// class Object
	obj := dc.ObjectBase()
	obj.GUID(ObjectGUID, Public)
	obj.Uint32(ObjectType, Public)
	obj.Uint32(ObjectEntry, Public)
	obj.Float32(ObjectScaleX, Public)
	obj.Pad()

	// class Item : Object
	item := obj.Extend("Item")
	item.GUID(ItemOwner, Public)
	item.GUID(ItemContained, Public)
	item.GUID(ItemCreator, Public)
	item.GUID(ItemGiftCreator, Public)
	item.Uint32(ItemStackCount, Public)
	item.Uint32(ItemDuration, Public)
	item.Uint32Array(ItemSpellCharges, 5, Public)
	item.Uint32(ItemFlags, Public)
	item.Uint32Array(ItemEnchantment, 21, Public)
	item.Uint32(ItemPropertySeed, Public)
	item.Uint32(ItemRandomPropertiesID, Public)
	item.Uint32(ItemTextID, Public)
	item.Uint32(ItemDurability, Public)
	item.Uint32(ItemMaxDurability, Public)

	// class Container : Item
	container := item.Extend("Container")
	container.Uint32(ContainerNumSlots, Public)
	container.Uint32(ContainerAlignPad, Public)
	container.GUIDArray(ContainerSlots, 36, Public)

	// class Unit : Object
	unit := obj.Extend("Unit")
	unit.GUID(UnitCharm, Public)
	unit.GUID(UnitSummon, Public)
	unit.GUID(UnitCharmedBy, Public)
	unit.GUID(UnitSummonedBy, Public)
	unit.GUID(UnitCreatedBy, Public)
	unit.GUID(UnitTarget, Public)
	unit.GUID(UnitPersuaded, Public)
	unit.GUID(UnitChannelObject, Public)
	unit.Uint32(UnitHealth, Public)
	unit.Uint32Array(UnitPowers, 5, Public)
	unit.Uint32(UnitMaxHealth, Public)
	unit.Uint32Array(UnitMaxPowers, 5, Public)
	unit.Uint32(UnitLevel, Public)
	unit.Uint32(UnitFactionTemplate, Public)

	unit.Byte(UnitRace, Public)
	unit.Byte(UnitClass, Public)
	unit.Byte(UnitGender, Public)
	unit.Byte(UnitPower, Public)

	unit.Uint32Array(UnitVirtualItemSlotIDs, 3, Public)
	unit.Uint32Array(UnitVirtualItemInfos, 6, Public)
	unit.Uint32(UnitFlags, Public)
	unit.Uint32Array(UnitAuras, 48, Public)
	unit.Uint32Array(UnitAuraFlags, 6, Public)
	unit.Uint32Array(UnitAuraLevels, 12, Public)
	unit.Uint32Array(UnitAuraApplications, 12, Public)
	unit.Uint32(UnitAuraState, Public)
	unit.Uint32(UnitBaseAttackTime, Public)
	unit.Uint32(UnitOffhandAttackTime, Public)
	unit.Uint32(UnitRangedAttackTime, Public)
	unit.Float32(UnitBoundingRadius, Public)
	unit.Float32(UnitCombatReach, Public)
	unit.Uint32(UnitDisplayID, Public)
	unit.Uint32(UnitNativeDisplayID, Public)
	unit.Uint32(UnitMountDisplayID, Public)
	unit.Float32(UnitMinDamage, Public)
	unit.Float32(UnitMaxDamage, Public)
	unit.Uint32(UnitMinOffhandDamage, Public)
	unit.Uint32(UnitMaxOffhandDamage, Public)

	unit.Byte(UnitStandState, Public)
	unit.Byte(UnitLoyaltyLevel, Public)
	unit.Byte(UnitShapeshiftForm, Public)
	unit.Byte(UnitStandMiscFlags, Public)

	unit.Uint32(UnitPetNumber, Public)
	unit.Uint32(UnitPetNameTimestamp, Public)
	unit.Uint32(UnitPetExperience, Public)
	unit.Uint32(UnitPetNextLevelExp, Public)
	unit.Uint32(UnitDynamicFlags, Public)
	unit.Uint32(UnitChannelSpell, Public)
	unit.Float32(UnitModCastSpeed, Public)
	unit.Uint32(UnitCreatedBySpell, Public)
	unit.Uint32(UnitNPCFlags, Public)
	unit.Uint32(UnitNPCEmoteState, Public)
	unit.Uint32(UnitTrainingPoints, Public)
	unit.Uint32Array(UnitStats, 5, Public)
	unit.Uint32Array(UnitResistances, 7, Public)
	unit.Uint32(UnitBaseMana, Public)
	unit.Uint32(UnitBaseHealth, Public)

	unit.Byte(UnitSheathState, Public)
	unit.Byte(UnitAuraByteFlags, Public)
	unit.Byte(UnitPetRename, Public)
	unit.Byte(UnitPetShapeshiftForm, Public)

	unit.Int32(UnitAttackPower, Public)
	unit.Int32(UnitAttackPowerMods, Public)
	unit.Float32(UnitAttackPowerMultiplier, Public)
	unit.Int32(UnitRangedAttackPower, Public)
	unit.Int32(UnitRangedAttackPowerMods, Public)
	unit.Float32(UnitRangedAttackPowerMultiplier, Public)
	unit.Float32(UnitMinRangedDamage, Public)
	unit.Float32(UnitMaxRangedDamage, Public)
	unit.Uint32Array(UnitPowerCostModifier, 7, Public)
	unit.Float32Array(UnitPowerCostMultiplier, 7, Public)
	unit.Pad()

	// class Player : Unit
	plyr := unit.Extend("Player")
	plyr.GUID(PlayerDuelArbiter, Public)
	plyr.Uint32(PlayerFlags, Public)
	plyr.Uint32(PlayerGuildID, Public)
	plyr.Uint32(PlayerGuildRank, Public)

	plyr.Byte(PlayerSkin, Public)
	plyr.Byte(PlayerFace, Public)
	plyr.Byte(PlayerHairStyle, Public)
	plyr.Byte(PlayerHairColor, Public)

	plyr.Byte(PlayerFacialHair, Public)
	plyr.Byte(PlayerRestBits, Public)
	plyr.Byte(PlayerBankBagSlotCount, Public)
	plyr.Byte(PlayerRestState, Public)

	plyr.Byte(PlayerGender, Public)
	plyr.Byte(PlayerGenderUnk, Public)
	plyr.Byte(PlayerDrunkness, Public)
	plyr.Byte(PlayerPVPRank, Public)

	plyr.Uint32(PlayerDuelTeam, Public)
	plyr.Uint32(PlayerGuildTimestamp, Public)

	questLog := plyr.Array(PlayerQuestLog, 20)
	questLog.Uint32("QuestID", Public)
	questLog.Uint32("CountState", Public)
	questLog.Uint32("Time", Public)
	questLog.End()

	visItems := plyr.Array(PlayerVisibleItems, 19)
	visItems.GUID("Creator", Public)
	visItems.Uint32("Entry", Public)
	visItems.Uint32Array("Enchantments", 7, Public)
	visItems.Uint32("Properties", Public)
	visItems.Pad()
	visItems.End()

	plyr.GUIDArray(PlayerInventorySlots, 39, Private)
	plyr.GUIDArray(PlayerBankSlots, 24, Private)
	plyr.GUIDArray(PlayerBankBagSlots, 6, Private)
	plyr.GUIDArray(PlayerVendorBuybackSlots, 12, Private)
	plyr.GUIDArray(PlayerKeyringSlots, 32, Private)
	plyr.GUID(PlayerFarSight, Public)
	plyr.GUID(PlayerFieldComboTarget, Public)
	plyr.Uint32(PlayerXP, Public)
	plyr.Uint32(PlayerNextLevelXP, Public)
	plyr.Uint32Array(PlayerSkillInfos, 384, Private)
	plyr.Uint32Array(PlayerCharacterPoints, 2, Private)
	plyr.Uint32(PlayerTrackCreatures, Private)
	plyr.Uint32(PlayerTrackResources, Private)
	plyr.Float32(PlayerBlockPercentage, Public)
	plyr.Float32(PlayerDodgePercentage, Public)
	plyr.Float32(PlayerParryPercentage, Public)
	plyr.Float32(PlayerCritPercentage, Public)
	plyr.Float32(PlayerRangedCritPercentage, Public)
	plyr.Uint32Array(PlayerExploredZones, 64, Public)
	plyr.Uint32(PlayerRestStateExperience, Public)
	plyr.Int32(PlayerCoinage, Private)
	plyr.Uint32Array(UnitPosStats, 5, Public)
	plyr.Uint32Array(UnitNegStats, 5, Public)
	plyr.Uint32Array(UnitResistanceBuffModsPositive, 7, Public)
	plyr.Uint32Array(UnitResistanceBuffModsNegative, 7, Public)
	plyr.Uint32Array(PlayerModDamageDonePositive, 7, Public)
	plyr.Uint32Array(PlayerModDamageDoneNegative, 7, Public)
	plyr.Float32Array(PlayerModDamageDonePercentage, 7, Public)

	plyr.Byte(PlayerFieldBytesFlags, Public)
	plyr.Byte(PlayerFieldBytesUnk1, Public)
	plyr.Byte(PlayerActionBarToggle, Public)
	plyr.Byte(PlayerFieldBytesUnk2, Public)

	plyr.Uint32(PlayerAmmoID, Public)
	plyr.Uint32(PlayerSelfResSpell, Public)
	plyr.Uint32(PlayerPVPMedals, Public)
	plyr.Uint32Array(PlayerBuybackPrices, 12, Private)
	plyr.Uint32Array(PlayerBuybackTimestamps, 12, Private)
	plyr.Uint32(PlayerKills, Public)
	plyr.Uint32(PlayerYesterdayKills, Public)
	plyr.Uint32(PlayerLastWeekKills, Public)
	plyr.Uint32(PlayerThisWeekKills, Public)
	plyr.Uint32(PlayerThisWeekContribution, Public)
	plyr.Uint32(PlayerLifetimeHonorableKills, Public)
	plyr.Uint32(PlayerLifetimeDishonorableKills, Public)
	plyr.Uint32(PlayerYesterdayContribution, Public)
	plyr.Uint32(PlayerLastWeekContribution, Public)
	plyr.Uint32(PlayerLastWeekRank, Public)

	plyr.Byte(PlayerHonorRankPoints, Public)
	plyr.Byte(PlayerDetectionFlags, Public)

	plyr.Int32(PlayerWatchedFactionIndex, Public)
	plyr.Uint32Array(PlayerCombatRatings, 20, Public)

	// class GameObject : Object
	gobj := obj.Extend("GameObject")

	gobj.GUID(GObjectCreatedBy, Public)
	gobj.Uint32(GObjectDisplayID, Public)
	gobj.Uint32(GObjectFlags, Public)
	gobj.Float32Array(GObjectRotation, 4, Public)
	gobj.Uint32(GObjectState, Public)
	gobj.Float32(GObjectPosX, Public)
	gobj.Float32(GObjectPosY, Public)
	gobj.Float32(GObjectPosZ, Public)
	gobj.Float32(GObjectFacing, Public)
	gobj.Uint32(GObjectDynamicFlags, Public)
	gobj.Uint32(GObjectFaction, Public)
	gobj.Uint32(GObjectTypeID, Public)
	gobj.Uint32(GObjectLevel, Public)
	gobj.Uint32(GObjectArtKit, Public)
	gobj.Uint32(GObjectAnimProgress, Public)
	gobj.Uint32(GObjectPadding, Public)

	// class DynamicObject : Object
	dobj := obj.Extend("DynamicObject")
	dobj.GUID(DynamicObjectCaster, Public)

	dobj.Byte(DynamicObjectType, Public)

	dobj.Uint32(DynamicObjectSpellID, Public)
	dobj.Float32(DynamicObjectRadius, Public)
	dobj.Float32(DynamicObjectPosX, Public)
	dobj.Float32(DynamicObjectPosY, Public)
	dobj.Float32(DynamicObjectPosZ, Public)
	dobj.Float32(DynamicObjectFacing, Public)

	// class Corpse : Object
	corp := obj.Extend("Corpse")
	corp.GUID(CorpseOwner, Public)
	corp.Float32(CorpseFacing, Public)
	corp.Float32(CorpsePosX, Public)
	corp.Float32(CorpsePosY, Public)
	corp.Float32(CorpsePosZ, Public)
	corp.Uint32(CorpseDisplayID, Public)
	corp.Uint32Array(CorpseItem, 19, Public)

	corp.Byte(CorpsePlayerUnk, Public)
	corp.Byte(CorpseRace, Public)
	corp.Byte(CorpseGender, Public)
	corp.Byte(CorpseSkin, Public)

	corp.Byte(CorpseFace, Public)
	corp.Byte(CorpseHairStyle, Public)
	corp.Byte(CorpseHairColor, Public)
	corp.Byte(CorpseFacialHair, Public)

	corp.Uint32(CorpseGuild, Public)
	corp.Uint32(CorpseFlags, Public)
	corp.Uint32(CorpseDynamicFlags, Public)
	corp.Pad()

	Descriptors[5875] = dc
}
