package worldserver

import (
	"io"
	"time"

	"github.com/superp00t/etc"
	"github.com/superp00t/etc/yo"
	"github.com/superp00t/gophercraft/gcore"
	"github.com/superp00t/gophercraft/guid"
	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/packet/update"
	"github.com/superp00t/gophercraft/worldserver/wdb"
)

func (s *Session) SetupOnLogin() {
	s.State = InWorld

	s.SendVerifyLoginPacket()
	s.SendAccountDataTimes()
	s.SendRestStart()
	s.BindpointUpdate()
	s.SendTutorialFlags()
	s.SendSpellList()
	s.SendActionButtons()
	s.SendFactions()
	s.SetTimeSpeed()
	s.SpawnPlayer()

	yo.Ok("Sent spawn packet.")

	go func() {
		time.Sleep(500 * time.Millisecond)
		s.Annf("Welcome to Gophercraft %s!", gcore.Version)
	}()
}

func packTime(t time.Time) int32 {
	year, month, day := t.Date()
	tm_year := int32(year)
	tm_mon := int32(month)
	tm_mday := int32(day)
	tm_wday := int32(t.Weekday())
	tm_hour := int32(t.Hour())
	tm_min := int32(t.Minute())

	return (tm_year-100)<<24 | tm_mon<<20 | (tm_mday-1)<<14 | tm_wday<<11 | tm_hour<<6 | tm_min
}

func (s *Session) SetTimeSpeed() {
	pkt := packet.NewWorldPacket(packet.SMSG_LOGIN_SETTIMESPEED)
	pkt.WriteInt32(packTime(time.Now()))
	pkt.WriteFloat32(0.01666667)

	s.SendAsync(pkt)
}

func (s *Session) SendVerifyLoginPacket() {
	v := packet.NewWorldPacket(packet.SMSG_LOGIN_VERIFY_WORLD)

	v.WriteUint32(s.Char.Map)
	v.WriteFloat32(s.Char.X)
	v.WriteFloat32(s.Char.Y)
	v.WriteFloat32(s.Char.Z)
	v.WriteFloat32(s.Char.O)

	s.SendAsync(v)
	yo.Ok("Sent verify login packet")
}

func (s *Session) SendRestStart() {
	v := packet.NewWorldPacket(packet.SMSG_QUEST_FORCE_REMOVE)
	v.WriteUint32(0)
	s.SendAsync(v)
}

func (s *Session) SendAccountDataTimes() {
	v := packet.NewWorldPacket(packet.SMSG_ACCOUNT_DATA_TIMES)
	for i := 0; i < 32; i++ {
		v.WriteUint32(0)
	}
	s.SendAsync(v)
}

func (s *Session) SendTutorialFlags() {
	v3 := packet.NewWorldPacket(packet.SMSG_TUTORIAL_FLAGS)
	for i := 0; i < 8; i++ {
		v3.WriteUint32(0xFFFFFFFF)
	}
	s.SendAsync(v3)
	yo.Println("Tutorial flags sent.")
}

func (s *Session) HandleAccountDataUpdate(data []byte) {
	yo.Spew(data)
}

func (s *Session) SendSpellList() {
	spl := []byte{0x0, 0x23, 0x0, 0xB, 0x56, 0x0, 0x0, 0x93, 0x54, 0x0, 0x0, 0x78, 0x50, 0x0, 0x0, 0x77, 0x50, 0x0, 0x0, 0x76, 0x50, 0x0, 0x0, 0x75, 0x50, 0x0, 0x0, 0xA5, 0x23, 0x0, 0x0, 0x76, 0x23, 0x0, 0x0, 0xC2, 0x20, 0x0, 0x0, 0x9C, 0x2, 0x0, 0x0, 0xCC, 0x0, 0x0, 0x0, 0x94, 0x54, 0x0, 0x0, 0x4E, 0x9, 0x0, 0x0, 0xAF, 0x9, 0x0, 0x0, 0x51, 0x0, 0x0, 0x0, 0x91, 0x13, 0x0, 0x0, 0xA8, 0x0, 0x0, 0x0, 0xA, 0x2, 0x0, 0x0, 0x1A, 0x59, 0x0, 0x0, 0xCB, 0x0, 0x0, 0x0, 0x66, 0x18, 0x0, 0x0, 0xEA, 0xB, 0x0, 0x0, 0x80, 0x51, 0x0, 0x0, 0x67, 0x18, 0x0, 0x0, 0x4D, 0x19, 0x0, 0x0, 0x9B, 0x13, 0x0, 0x0, 0xE3, 0x0, 0x0, 0x0, 0x85, 0x0, 0x0, 0x0, 0x4E, 0x19, 0x0, 0x0, 0xBB, 0x1C, 0x0, 0x0, 0xCB, 0x19, 0x0, 0x0, 0x25, 0xD, 0x0, 0x0, 0x62, 0x1C, 0x0, 0x0, 0x59, 0x18, 0x0, 0x0, 0x63, 0x1C, 0x0, 0x0, 0x0, 0x0}

	p := packet.NewWorldPacket(packet.SMSG_INITIAL_SPELLS)
	p.Buffer = etc.FromBytes(spl)
	s.SendAsync(p)
}

func (s *Session) SendActionButtons() {
	Buf_actionbuttons := []byte{0xCB, 0x19, 0x0, 0x0, 0x85, 0x0, 0x0, 0x0, 0xA8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9F, 0x0, 0x0, 0x80, 0x16, 0x8, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	p := packet.NewWorldPacket(packet.SMSG_ACTION_BUTTONS)
	p.Buffer = etc.FromBytes(Buf_actionbuttons)
	s.SendAsync(p)
}

func (s *Session) SendFactions() {
	Buf_reps := []byte{0x40, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x9, 0x0, 0x0, 0x0, 0x0, 0xE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x11, 0x0, 0x0, 0x0, 0x0, 0x11, 0x0, 0x0, 0x0, 0x0, 0x11, 0x0, 0x0, 0x0, 0x0, 0x11, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x4, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x14, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x0, 0x18, 0x0, 0x0, 0x0, 0x0, 0xE, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x10, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	p := packet.NewWorldPacket(packet.SMSG_INITIALIZE_FACTIONS)
	p.Buffer = etc.FromBytes(Buf_reps)
	s.SendAsync(p)
}

func (s *Session) GetPlayerRace() packet.Race {
	return packet.Race(s.GetByteValue(update.UnitRace))
}

func (s *Session) GetPlayerLevel() uint8 {
	return uint8(s.GetUint32Value(update.UnitLevel))
}

// SpawnPlayer initializes the player into the object manager and sends the packets needed to log the client into the world.
func (s *Session) SpawnPlayer() {
	s.WS.PlayersL.Lock()
	s.WS.PlayerList[s.PlayerName()] = s
	s.WS.PlayersL.Unlock()

	// fill out attribute fields
	s.PlayerSpeeds = update.Speeds{
		update.Walk:         2.5,
		update.Run:          7,
		update.RunBackward:  4.5,
		update.Swim:         4.722222,
		update.SwimBackward: 2.5,
		update.Turn:         3.141594,
	}

	s.ValuesBlock = update.NewValuesBlock()
	s.SetGUIDValue(update.ObjectGUID, s.GUID())
	s.SetTypeMask(s.Version(), guid.TypeMaskObject|guid.TypeMaskUnit|guid.TypeMaskPlayer)
	s.SetFloat32Value(update.ObjectScaleX, 1.0)

	s.SetUint32Value(update.UnitHealth, 80)
	s.SetUint32ArrayValue(
		update.UnitPowers,
		4143,
		nil,
		nil,
		100,
		nil,
	)
	s.SetUint32Value(update.UnitMaxHealth, 80)
	s.SetUint32ArrayValue(
		update.UnitMaxPowers,
		4143,
		1000,
		nil,
		100,
		nil,
	)
	s.SetUint32Value(update.UnitLevel, uint32(s.Char.Level))
	s.SetUint32Value(update.UnitFactionTemplate, 1)

	s.SetByteValue(update.UnitRace, uint8(s.Char.Race))
	s.SetByteValue(update.UnitClass, uint8(s.Char.Class))
	s.SetByteValue(update.UnitGender, uint8(s.Char.Gender))
	s.SetByteValue(update.UnitPower, PowerType(packet.Class(s.Char.Class)))

	// TODO: replace with SetBitValue()
	s.SetUint32Value(update.UnitFlags, 8)
	s.SetUint32Value(update.PlayerFlags, 0x0020)

	s.SetUint32Value(update.UnitBaseAttackTime, 2900)
	s.SetUint32Value(update.UnitOffhandAttackTime, 2000)

	s.SetFloat32Value(update.UnitBoundingRadius, 1.0)
	s.SetFloat32Value(update.UnitCombatReach, 1.0)
	s.SetUint32Value(update.UnitDisplayID, s.WS.GetNative(packet.Race(s.Char.Race), s.Char.Gender))
	s.SetUint32Value(update.UnitNativeDisplayID, s.WS.GetNative(packet.Race(s.Char.Race), s.Char.Gender))
	s.SetUint32Value(update.UnitMountDisplayID, 0)

	s.SetFloat32Value(update.UnitMinDamage, 50)
	s.SetFloat32Value(update.UnitMaxDamage, 50)
	s.SetUint32Value(update.UnitMinOffhandDamage, 50)
	s.SetUint32Value(update.UnitMaxOffhandDamage, 50)

	s.SetByteValue(update.UnitLoyaltyLevel, 0xEE)

	s.SetFloat32Value(update.UnitModCastSpeed, 30)
	s.SetUint32ArrayValue(
		update.UnitStats,
		0,
		0,
		0,
		0,
		0,
	)

	s.SetUint32ArrayValue(
		update.UnitResistances,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	)

	s.SetUint32Value(update.UnitBaseMana, 60)
	// todo: replace with bit fields
	// s.SetBitValue(update.UnitAuraByteFlagSupportable, true)
	// s.SetBitValue(update.UnitAuraByteFlagNoDispel, true)
	s.SetByteValue(update.UnitAuraByteFlags, 0x08|0x20)

	s.SetInt32Value(update.UnitAttackPower, 20)
	s.SetInt32Value(update.UnitAttackPowerMods, 0)

	s.SetInt32Value(update.UnitRangedAttackPower, 0)
	s.SetInt32Value(update.UnitRangedAttackPowerMods, 0)

	s.SetFloat32Value(update.UnitMinRangedDamage, 0)
	s.SetFloat32Value(update.UnitMaxRangedDamage, 0)

	s.SetByteValue(update.PlayerSkin, s.Char.Skin)
	s.SetByteValue(update.PlayerFace, s.Char.Face)
	s.SetByteValue(update.PlayerHairStyle, s.Char.HairStyle)
	s.SetByteValue(update.PlayerHairColor, s.Char.HairColor)

	s.SetByteValue(update.PlayerFacialHair, s.Char.FacialHair)
	s.SetByteValue(update.PlayerBankBagSlotCount, 8)
	s.SetByteValue(update.PlayerRestState, 0x01)

	s.SetByteValue(update.PlayerGender, s.Char.Gender)

	s.SetUint32Value(update.PlayerXP, 0)
	s.SetUint32Value(update.PlayerNextLevelXP, 2500)

	s.SetUint32ArrayValue(update.PlayerCharacterPoints, 51, 2)

	s.SetFloat32Value(update.PlayerBlockPercentage, 4.0)
	s.SetFloat32Value(update.PlayerDodgePercentage, 4.0)
	s.SetFloat32Value(update.PlayerParryPercentage, 4.0)
	s.SetFloat32Value(update.PlayerCritPercentage, 4.0)

	s.SetUint32Value(update.PlayerRestStateExperience, 200)
	s.SetInt32Value(update.PlayerCoinage, 50000)

	s.PlayerPosition = update.Quaternion{
		Point3: update.Point3{
			X: s.Char.X,
			Y: s.Char.Y,
			Z: s.Char.Z,
		},
		O: s.Char.O,
	}

	s.CurrentMap = s.Char.Map

	// send player create packet of themself
	s.SendObjectCreate(s)

	cMap := s.Map()

	// add our player to map, and notify nearby players of their presence
	cMap.AddObject(s)

	// notify our player of these nearby players.
	for _, currentPlayer := range cMap.NearbySessions(s) {
		s.SendObjectCreate(currentPlayer)
	}
}

func (s *Session) BindpointUpdate() {
	//goldshire
	p := packet.NewWorldPacket(packet.SMSG_BINDPOINTUPDATE)
	p.WriteFloat32(-8949.95)
	p.WriteFloat32(-132.493)
	p.WriteFloat32(83.5312)
	p.WriteUint32(0)
	p.WriteUint32(12)

	s.SendAsync(p)
}

func (s *Session) HandlePlayernameQuery(e *etc.Buffer) {
	g, err := guid.DecodeUnpacked(s.Version(), e)
	if err != nil {
		panic(err)
	}

	var chars []wdb.Character

	s.WS.DB.Where("id = ?", g.Counter()).Find(&chars)
	if len(chars) == 0 {
		yo.Warn("No such data exists for", g)
		return
	}

	c := chars[0]

	resp := packet.NewWorldPacket(packet.SMSG_NAME_QUERY_RESPONSE)
	resp.WriteUint64(g.Classic())
	resp.WriteCString(c.Name)
	resp.WriteByte(0)
	resp.WriteUint32(uint32(c.Race))
	resp.WriteUint32(uint32(c.Gender))
	resp.WriteUint32(uint32(c.Class))
	s.SendAsync(resp)
}

func (s *Session) encodePackedGUID(wr io.Writer, g guid.GUID) {
	g.EncodePacked(s.Version(), wr)
}

func (s *Session) HandleMoves(t packet.WorldType, b []byte) {
	e, err := update.DecodeMovementInfo(s.Version(), etc.FromBytes(b))
	if err != nil {
		yo.Warn(err)
		return
	}

	s.PlayerPosition = e.Position

	for _, v := range s.Map().NearbySessions(s) {
		// yo.Ok("Relaying moves", t, s.Char.Name, "->", v.Char.Name)
		out := packet.NewWorldPacket(t)
		s.encodePackedGUID(out, s.GUID())
		update.EncodeMovementInfo(v.Version(), out.Buffer, e)
		v.SendAsync(out)
	}
}

func (s *Session) HandleLogoutRequest(b []byte) {
	// TODO: deny if in combat

	resp := packet.NewWorldPacket(packet.SMSG_LOGOUT_RESPONSE)
	resp.WriteUint32(0)
	resp.WriteByte(0)

	s.SendAsync(resp)

	s.State = CharacterSelectMenu

	s.Map().RemoveObject(s.GUID())

	resp = packet.NewWorldPacket(packet.SMSG_LOGOUT_COMPLETE)
	s.SendAsync(resp)
}
