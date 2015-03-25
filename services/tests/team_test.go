package services

import (
	. "github.com/panyam/backbone/services/core"
	. "gopkg.in/check.v1"
	// "code.google.com/p/gomock/gomock"
	"log"
	// "time"
)

func (s *TestSuite) TestCreateTeam(c *C) {
	svc := s.serviceGroup.TeamService
	team, err := svc.CreateTeam(0, "group", "test")
	log.Println("Error: ", err)
	c.Assert(err, Equals, nil)
	c.Assert(team, Not(Equals), (*Team)(nil))
	c.Assert(team.Name, Equals, "test")
	team, err = svc.GetTeamByName("group", "test")
	c.Assert(err, Equals, nil)
	c.Assert(team, Not(Equals), (*Team)(nil))
	c.Assert(team.Name, Equals, "test")
}

/*
func (s *TestSuite) TestCreateTeamExistsByName(c *C) {
	svc := s.serviceGroup.TeamService
	team, err := svc.CreateTeam("1", "group", "test")
	team, err = svc.CreateTeam("1", "group2", "test2")
	c.Assert(err, Not(Equals), nil)
	c.Assert(team, Equals, (*Team)(nil))
}

func (s *TestSuite) TestDeleteTeam(c *C) {
	svc := s.serviceGroup.TeamService
	team, err := svc.CreateTeam(0, "group", "test")
	c.Assert(err, Equals, nil)
	c.Assert(team, Not(Equals), (*Team)(nil))
	svc.DeleteTeam(team)
	team, err = svc.GetTeamByName("group", "test")
	c.Assert(err, Not(Equals), nil)
	c.Assert(team, Equals, (*Team)(nil))
}

func (s *TestSuite) TestJoinTeam(c *C) {
	svc := s.serviceGroup.TeamService
	team, _ := svc.CreateTeam(0, "group", "test")
	user := NewUser("1", "user1")
	svc.JoinTeam(team, user)
	c.Assert(svc.TeamContains(team, user), Equals, true)
}

func (s *TestSuite) TestLeaveTeam(c *C) {
	svc := s.serviceGroup.TeamService
	team, _ := svc.CreateTeam("", "group", "test")
	user := NewUser("1", "user1")
	svc.JoinTeam(team, user)
	c.Assert(svc.TeamContains(team, user), Equals, true)
	svc.LeaveTeam(team, user)
	c.Assert(svc.TeamContains(team, user), Equals, false)
}
*/
