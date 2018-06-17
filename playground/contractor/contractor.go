package main

import "fmt"

// Board represents a surface we can work on
type Board struct {
	NailDriven int
	NailNeeded int
}

//NailDriver represents behavior to drive nail into a board
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

//NailPuller represents behavior to remove nail into a board
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

//NailDrivPuller represents behavior to drive/remove nail into a board
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

//Mallet is a tool that pounds in nails
type Mallet struct {

}

//DriveNail pounds a nail into the specific Board
func (Mallet)DriveNail(nailSupply *int, b *Board){
	//take a nail out of the supply
	*nailSupply--

	//pound a nail into the board
	b.NailDriven++

	fmt.Println("Mallet: Pound a nail into board")

}

// Crowbar is a tool that remove nails
type Crowbar struct {

}

//PullNail yanks a nail out of the special board
func (Crowbar)PullNail (nailsupply *int, b *Board){
	//Yank a nail out of board
	b.NailDriven--

	//put that nail back into the supply
	*nailsupply++

	fmt.Println("Crowbar:yanked a nail out of the board")
}

//contractor carries out the task of securing board.
type Contractor struct {

}

//Fasten will drive nails into a board.
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board){
	for(b.NailDriven < b.NailNeeded){
		d.DriveNail(nailSupply, b)
	}
}

//Unfasten will remove nails from a board.
func (Contractor)Unfasten(p NailPuller,nailSupply *int, b *Board){
	for(b.NailNeeded < b.NailDriven){
		p.PullNail(nailSupply,b)
	}
}

//ProcessBoard works against board.
func (c Contractor)ProcessBoard (dp NailDrivePuller, nailSupply *int, boards []Board){
	for i := range boards{
		b := &boards[i]

		fmt.Printf("contractor:examming board #%d: %+v\n",i+1,b)

		switch  {
		case b.NailDriven < b.NailNeeded:
			c.Fasten(dp, nailSupply, b)
		case b.NailNeeded < b.NailDriven:
			c.Unfasten(dp, nailSupply, b)

		}
	}
}

// Toolbox can contain any number of tools.
type Toolbox struct {
	NailPuller
	NailDriver
	nails int
}

func main(){
	//inventory the old boards to remove, and the new board will replace them.
	boards := []Board{
		//Robbot board to be removed.
		{NailDriven:3},
		{NailDriven:1},
		{NailDriven:6},

		//Fresh board to be fasten.
		{NailNeeded:2},
		{NailNeeded:5},
		{NailNeeded:4},

	}

	//fill a toolbox
	tb := Toolbox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails: 10,
	}
	//display the current state of our toolbox and state
	displayState(&tb,boards)

	//Hire a contractor and put our contractor to work
	var c Contractor
	c.ProcessBoard(&tb, &tb.nails,boards)

	//display the new state of our toolbox and boards.
	displayState(&tb, boards)


}

//displayState provide information about all boads.
func displayState(tb *Toolbox,boards []Board){
	fmt.Printf("Box: %#v\n",tb)
	fmt.Println("Boards:")

	for _,b := range boards{
		fmt.Printf("\t%+v\n",b)

	}
	fmt.Println()
}