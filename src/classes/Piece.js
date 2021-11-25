export default class Piece {
	constructor(org, final) {
		this.org = org
		this.final = final
		this.piece = org['piece']
		this.xo = org['loc'][1]
		this.xf = final['loc'][1]
		this.xd = this.xf - this.xo
		this.yo = org['loc'][0]
		this.yf = final['loc'][0]
		this.yd = this.yf - this.yo
	}

	getMovementRules() {
		if ( this.org['player'] == this.final['player'] ) {
			console.log("same player piece")
			return this.valid = false
		}
		else {
		switch (this.piece) {
			case "knight":
				this.knightMovement()
				break;
			case "rook":
				this.rookMovement()
				break;
			case "bishop":
				this.bishopMovement()
				break;
			case "queen":
				this.queenMovement()
				break;
			case "king":
				this.kingMovement()
				break;
			case "pawn":
				this.pawnMovement()
				break;
			default:
			console.log("Somethinga wenta wronga")
			return this.valid = false;
		}
		return this.valid
		}
	}

	invalid() {
		console.log('invalid move')
		return this.valid = false
	}

	knightMovement() {
		if (Math.abs(this.xd)+Math.abs(this.yd)==3 && Math.abs(Math.abs(this.xd)-Math.abs(this.yd))==1) {
			return this.valid = true
		}
		this.invalid()
	}

	bishopMovement() {
		if (Math.abs(this.xd) == Math.abs(this.yd)) {
			return this.valid = true
		}
		this.invalid()
	}

	rookMovement() {
		if ((this.xd == 0 && this.yd != 0) || (this.yd == 0 && this.xd != 0)) {
			return this.valid = true
		}
		this.invalid()
	}

	queenMovement() {
		if (this.bishopMovement() || this.rookMovement()) {
			return this.valid = true
		}
		this.invalid()
	}

	kingMovement() {
		if ((Math.abs(this.xd) == 1 || Math.abs(this.yd) == 1) && (this.bishopMovement() || this.rookMovement())) {
			return this.valid = true
		}
		this.invalid()
	}

	pawnMovement() {
		if ((this.org['player'] == 'b' && ((this.yd == 1 && this.xd == 0) || (this.yo == 2 && this.yd == 2))) || (this.org['player'] == 'w' && ((this.yd == -1 && this.xd == 0) || (this.yo == 7 && this.yd == -2)))) {
			return this.valid = true
		}
		this.invalid()
	}

}
