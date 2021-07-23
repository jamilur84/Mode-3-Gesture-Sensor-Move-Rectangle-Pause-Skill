package MultiHexaLeapMotSkill

import (

	//"math"
	"os"
	"time"

	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
)

const (
	TIME_MOVE_DURATION    	= 2000 // milliseconds
	TIME_MOVE_DURATION_LONG = 5000 // milliseconds
	TIME_PAUSE_MIDDLE    	= 5000 // milliseconds
	MOVE_HEAD_DURATION      = 500  // milliseconds
	MOVE_LEG_DURATION       = 1000 // milliseconds
	WALK_SPEED              = 1.5  // cm per second
	MOVE_LEFT               = 90   // Degrees to the left
	MOVE_RIGHT              = 270  // Degress to the right
	MOVE_FWD                = 0	 // Degress to the front 
	MOVE_BACKWD             = 180  // Degress to the back
	HEIGHT_STAND_UP         = 50   // The height for the stand-up movement
	HEIGHT_STAND_DOWN       = -10  // The height for the stand-down movement
	HEAD_POS_0_CAL     	    = 0    // The default head position
	LEFT_LEG			    = 1    // The left leg index
	RIGHT_LEG			    = 0    // The right leg index
	MAX_CNTR_WALK_RECT	    = 3	 // The maximum value of the counter for walk rectangle behavior
)

// The data structure for the new skill
type MultiHexaLeapMotSkill struct {
	skill.Base
	stop      chan bool
	direction float64
	is_busy   bool
	cntr_walk_rect int
}

// The function for registering the new skill
func NewSkill() skill.Interface {
	return &MultiHexaLeapMotSkill{
		stop: make(chan bool),
	}
}

// Move the robot to the stand-up position
func (d *MultiHexaLeapMotSkill) standup() {
	// Set the is busy flag to true
	d.is_busy = true
	// Set the robot height to the stand-up position
	hexabody.StandWithHeight(HEIGHT_STAND_UP)
	
	// Reset the is busy flag
	d.is_busy = false
}

// Move the robot to the stand-down position
func (d *MultiHexaLeapMotSkill) standdown() {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Set the robot height to the stand-down position
	hexabody.StandWithHeight(HEIGHT_STAND_DOWN)
	
	// Reset the is busy flag
	d.is_busy = false
}

// Move the robot in a rectangle path
func (d *MultiHexaLeapMotSkill) walkPathRectangle(increment int) {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Increment or decrement the multiplier for walking distance
	if increment > 0 {
		if d.cntr_walk_rect < MAX_CNTR_WALK_RECT {
			d.cntr_walk_rect = d.cntr_walk_rect + 1
		}
	} else if increment < 0 {
		if d.cntr_walk_rect > 1 {
			d.cntr_walk_rect = d.cntr_walk_rect - 1
		}
	}
	
	// The full duration of the robot movement
	var full_dur time.Duration = time.Duration(TIME_MOVE_DURATION_LONG) * time.Millisecond * time.Duration(d.cntr_walk_rect)
	
	// Walk straight ahead for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time. It is the half of time as the other path.
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_RIGHT, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_BACKWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_LEFT, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// Reset the is busy flag
	d.is_busy = false
}


// Move the robot in a rectangle path. It is the second version.
// The walking direction is determined by the head orientation.
func (d *MultiHexaLeapMotSkill) walkPathRectangle2(increment int) {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Increment or decrement the multiplier for walking distance
	if increment > 0 {
		if d.cntr_walk_rect < MAX_CNTR_WALK_RECT {
			d.cntr_walk_rect = d.cntr_walk_rect + 1
		}
	} else if increment < 0 {
		if d.cntr_walk_rect > 1 {
			d.cntr_walk_rect = d.cntr_walk_rect - 1
		}
	}
	
	// The full duration of the robot movement
	var full_dur time.Duration = time.Duration(TIME_MOVE_DURATION_LONG) * time.Millisecond * time.Duration(d.cntr_walk_rect)
	
	// STEP #1
	// Walk straight ahead for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time. It is the half of time as the other path.
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// STEP #2
	// Move the head 90 degrees to the right to change the direction
	hexabody.MoveHead(270, MOVE_HEAD_DURATION)
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// STEP #3
	// Move the head 90 degrees to the right to change the direction
	hexabody.MoveHead(180, MOVE_HEAD_DURATION)
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// STEP #4
	// Move the head 90 degrees to the right to change the direction
	hexabody.MoveHead(90, MOVE_HEAD_DURATION)
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()

	// Move the head back to the origin
	hexabody.MoveHead(HEAD_POS_0_CAL, MOVE_HEAD_DURATION)
	
	// Reset the is busy flag
	d.is_busy = false
}


// Move the robot in a rectangle path and pause in between.
// The walking direction is determined by the head orientation.
func (d *MultiHexaLeapMotSkill) walkPathRectPause(increment int) {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Increment or decrement the multiplier for walking distance
	if increment > 0 {
		if d.cntr_walk_rect < MAX_CNTR_WALK_RECT {
			d.cntr_walk_rect = d.cntr_walk_rect + 1
		}
	} else if increment < 0 {
		if d.cntr_walk_rect > 1 {
			d.cntr_walk_rect = d.cntr_walk_rect - 1
		}
	}
	
	// The full duration of the robot movement
	var full_dur time.Duration = time.Duration(TIME_MOVE_DURATION_LONG) * time.Millisecond * time.Duration(d.cntr_walk_rect)
	
	// STEP #1
	// Walk straight ahead for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time. It is the half of time as the other path.
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	// Pause before the next move.
	time.Sleep(TIME_PAUSE_MIDDLE)
	
	// STEP #2
	// Move the head 90 degrees to the right to change the direction
	hexabody.MoveHead(270, MOVE_HEAD_DURATION)
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	// Pause before the next move.
	time.Sleep(TIME_PAUSE_MIDDLE)
	
	// STEP #3
	// Move the head 90 degrees to the right to change the direction
	hexabody.MoveHead(180, MOVE_HEAD_DURATION)
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	// Pause before the next move.
	time.Sleep(TIME_PAUSE_MIDDLE)
	
	// STEP #4
	// Move the head 90 degrees to the right to change the direction
	hexabody.MoveHead(90, MOVE_HEAD_DURATION)
	// Walk straight to the right for a certain time
	hexabody.WalkContinuously(MOVE_FWD, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(full_dur)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	// Pause before the next move.
	time.Sleep(TIME_PAUSE_MIDDLE)

	// Move the head back to the origin
	hexabody.MoveHead(HEAD_POS_0_CAL, MOVE_HEAD_DURATION)
	
	// Reset the is busy flag
	d.is_busy = false
}

// Move the robot in a circle path
func (d *MultiHexaLeapMotSkill) walkPathCircle() {
	// Set the is busy flag to true
	d.is_busy = true
	
	// TODO:
	
	// Reset the is busy flag
	d.is_busy = false
}

// Move one of the robot's legs
func (d *MultiHexaLeapMotSkill) moveLeg(leg_idx int) {
	// Set the is busy flag to true
	d.is_busy = true
	
	var leg_pos_x float64 = 0
	
	if leg_idx == LEFT_LEG {
		leg_pos_x = 75.0
	} else if leg_idx == RIGHT_LEG {
		leg_pos_x = 75.0
	}
	
	// Raise the leg
	hexabody.MoveLeg(leg_idx, hexabody.NewLegPosition().SetCoordinates(leg_pos_x, 161.0, -89.0), MOVE_LEG_DURATION)
	
	// Wait for a certain period of time
	time.Sleep(MOVE_LEG_DURATION * time.Millisecond)
	
	// Bring back the leg to the original position
	hexabody.MoveLeg(leg_idx, hexabody.NewLegPosition().SetCoordinates(90, 122.0, 70.0), MOVE_LEG_DURATION)
	
	// Reset the is busy flag
	d.is_busy = false
}

// Tell the robot to walk to a certain direction for a certain period of time
func (d *MultiHexaLeapMotSkill) walk(direction float64) {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Align the head to the origin angle
	hexabody.MoveHead(HEAD_POS_0_CAL, MOVE_HEAD_DURATION)
	
	// Tell the Hexa to walk continuously to the given direction and speed 
	hexabody.WalkContinuously(direction, WALK_SPEED)
	// Wait for a certain period of time
	time.Sleep(TIME_MOVE_DURATION * time.Millisecond)
	// Stop the Hexa from walking
	hexabody.StopWalkingContinuously()
	
	// Reset the is busy flag
	d.is_busy = false
}

// Start the skill Mode 1.
// It tells the Hexa to do a set of movement to indicate the start of the interaction in Mode 1.
func (d *MultiHexaLeapMotSkill) startSkill1() {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Align the head to the origin angle
	hexabody.MoveHead(HEAD_POS_0_CAL, MOVE_HEAD_DURATION)
	
	// Pitch to the Left/Right at the start
	hexabody.Pitch(-20, 750)
	hexabody.Pitch(20, 1500)
	
	// Stand down and stand-up
	hexabody.StandWithHeight(HEIGHT_STAND_DOWN)
	hexabody.StandWithHeight(HEIGHT_STAND_UP)
	
	// Reset the is busy flag
	d.is_busy = false
}

// Start the skill Mode 2.
// It tells the Hexa to do a set of movement to indicate the start of the interaction in Mode 2.
func (d *MultiHexaLeapMotSkill) startSkill2() {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Raise up and down the front left leg
	d.moveLeg(LEFT_LEG)
	
	// Raise up and down the front right leg
	d.moveLeg(RIGHT_LEG)
	
	// Reset the is busy flag
	d.is_busy = false
}

// Start the skill Mode 3.
// It tells the Hexa to do a set of movement to indicate the start of the interaction in Mode 3.
func (d *MultiHexaLeapMotSkill) startSkill3() {
	// Set the is busy flag to true
	d.is_busy = true
	
	// Move the head 90 deg to the left
	hexabody.MoveHead(90, MOVE_HEAD_DURATION)
	// Move the head 90 deg to the right
	hexabody.MoveHead(270, MOVE_HEAD_DURATION)
	// Move the head back to the origin
	hexabody.MoveHead(HEAD_POS_0_CAL, MOVE_HEAD_DURATION)
	
	// Reset the is busy flag
	d.is_busy = false
}

// The OnStart event handler 
func (d *MultiHexaLeapMotSkill) OnStart() {
	// Initialize the counter for the walk rectangle behavior
	d.cntr_walk_rect = 1
	// Initialize the is busy flag
	d.is_busy = false
	// Use this method to do something when this skill is starting.
	log.Info.Println("MultiHexaLeapMotSkill Started")
}

// The OnClose event handler
func (d *MultiHexaLeapMotSkill) OnClose() {
	// Use this method to do something when this skill is closing.
	hexabody.Close()
}

// The OnConnect event handler
func (d *MultiHexaLeapMotSkill) OnConnect() {
	// Use this method to do something when the remote connected.
	err := hexabody.Start()
	
	if err != nil {
		log.Error.Println("Hexabody start err:", err)
		return
	}
}

// The OnDisconnect event handler
func (d *MultiHexaLeapMotSkill) OnDisconnect() {
	// Use this method to do something when the remote disconnected.
	os.Exit(0) // Closes the process when remote disconnects
}

// The OnRecvJSON event handler
func (d *MultiHexaLeapMotSkill) OnRecvJSON(data []byte) {
	// Use this method to do something when skill receive json data from remote client.
}

// The OnRecvString event handler
func (d *MultiHexaLeapMotSkill) OnRecvString(data string) {
	
	if d.is_busy == false {
		// Use this method to do something when skill receive string from remote client.
		switch data {
			case "start_mode1":
				go d.startSkill1()
			case "start_mode2":
				go d.startSkill2()
			case "start_mode3":
				go d.startSkill3()
			case "stop":
				// Reset the counter for the walk rectangle behavior
				d.cntr_walk_rect = 1
				// Set the stop flag
				d.stop <- true
				// Reset the is busy flag
				d.is_busy = false
				hexabody.MoveHead(HEAD_POS_0_CAL, MOVE_HEAD_DURATION)
				hexabody.StopWalkingContinuously()
				hexabody.Relax()
			case "left":
				go d.walk(MOVE_LEFT)
			case "right":
				go d.walk(MOVE_RIGHT)
			case "forward":
				go d.walk(MOVE_FWD)
			case "backward":
				go d.walk(MOVE_BACKWD)
			case "stand-up":
				go d.standup()
			case "stand-down":
				go d.standdown()
			case "gocircle":
				go d.walkPathCircle()
			case "gorect_p":
				go d.walkPathRectangle(1)
			case "gorect_n":
				go d.walkPathRectangle(-1)
			case "gorect2_p":
				go d.walkPathRectangle2(1)
			case "gorect2_n":
				go d.walkPathRectangle2(-1)
			case "gorectp_p":
				go d.walkPathRectPause(1)
			case "gorectp_n":
				go d.walkPathRectPause(-1)
			case "point-left":
				go d.moveLeg(LEFT_LEG)
			case "point-right":
				go d.moveLeg(RIGHT_LEG)
		}
	}
}
