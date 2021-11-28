import Piece from './Piece.js'

export default class Board {
	constructor() {
		self = this
		this.pieceLocation = [{}]
		this.totalTiles = this.buildBoard()
		this.selectedTile = 65
		this.currentTurn = "w"
	}

	setPieceLoc(pl) {
		this.pieceLocation = pl
		return this, self
	}

	changeTurn() {
		if (this.currentTurn == "w") {
			this.currentTurn = "b"
			console.log("turn is black")
			return this
		}
		else if (this.currentTurn == "b") {
			this.currentTurn = "w"
			console.log("turn is white")
			return this
		}
		else {
			console.log("Turn has error")
		}
	}

	buildBoard() {
		this.totalTiles = [...Array(9).keys()]
		self.totalTiles.shift()
		return this.totalTiles
	}

	getIndex(x, y) {
		var location = (x-1)*8+y-1
		return location
	}

	getPiece(x, y, pl) {
		var loc = self.getIndex(x,y)
		if (pl[loc] != undefined) {
		console.log(pl[loc])
		var image="https://chessapp.s3.us-west-2.amazonaws.com/" + pl[loc]['player'] + pl[loc]['piece'] + ".png"
		return image
		}
		console.log("page loading")
	}

	selected(x, y, pl) {
		var loc = self.getIndex(x,y)
		if (this.selectedTile != 65 && self.isValid(this.selectedTile, pl, loc)) {
			var placeHolder = this.selectedTile
			const sleep = (delay) => new Promise((resolve) => setTimeout(resolve, delay))
			const waitForAnimation = async () => {
			self.animatePiece(this.selectedTile,loc)
			await sleep(750)
			this.selectedTile = placeHolder
			document.getElementById('img'+(this.selectedTile)).style.transform=null
			pl[loc]['piece']=pl[this.selectedTile]['piece']
			pl[loc]['player']=pl[this.selectedTile]['player']
			pl[this.selectedTile]['piece']='none'
			pl[this.selectedTile]['player']=''
			}
			waitForAnimation()
			self.changeTurn()
			return
		}
		if (pl[loc]['piece']!='none') {
			this.selectedTile = loc
			return this.selectedTile
		}
	}

	isValid(select, pl, loc) {
		if (select != null && loc != select && new Piece(pl[select], pl[loc]).getMovementRules() == true && pl[select]['player'] == this.currentTurn) {
			if (pl[select]["piece"] == "knight" || self.checkPath(pl[select], pl[loc]) == true )
			return true
		}
		return false
	}

	animatePiece(st, loc) {
		var position = document.getElementById('tile'+(loc)).getBoundingClientRect()
		var f = document.getElementById('img'+(st))
		f.style.transform = 'translateY('+(position['y']-f.getBoundingClientRect()['y']+1)+'px)'
		f.style.transform += 'translateX('+(position['x']-f.getBoundingClientRect()['x']+2.5)+'px)'
	}

	checkPath(org,final) {
		var xo = org['loc'][0]
		var yo = org['loc'][1]
		var up = xo - final['loc'][0] > 0
		var valid = true
		var right = final['loc'][1] - yo > 0
		var udcenter = xo - final['loc'][0] == 0
		var lrcenter = final['loc'][1] - yo == 0
		var first = true
			while ((xo == final['loc'][0] && yo == final['loc'][1]) == false) {
				var index=self.getIndex(xo,yo)
				if (up == true && !udcenter) {
					xo--
				}
				if (up == false && !udcenter) {
					xo++
				}
				if (right == true && !lrcenter) {
					yo++
				}
				if (right == false && !lrcenter) {
					yo--
				}
				if (this.pieceLocation[index]["piece"] != "none" && first == false) {
					console.log("piece is in the way")
					return valid = false
				}
				first = false
			}
		return valid
	}
}
