<template>
 <div>
  <h2> {{ turnText }} </h2>
  <div class="boardBorder">
    <table class="board">
      <tr v-for="row in totalTiles" v-bind:key="row" class="row">
          <td v-for="column in totalTiles"  v-bind:key="column" v-bind:id="'tile'+getIndex(row, column)" @click="selectedTile=selected(row,column,pieceLocation,selectedTile)" class= "column">
            <div :class="{'active': getIndex(row, column) == selectedTile}">
                <img v-if="pieceLocation[getIndex(row, column)]" :src="getPiece(row,column,pieceLocation)" v-bind:id="'img'+getIndex(row, column)">
            </div>
          </td>
      </tr>
    </table>
  </div>
</div>
</template> 

<script>
import Board from "../classes/Board.js";
const waitForBoard = async () => {
const response = await fetch('https://www.maxelliotmills.com/board');
const boardJson = await response.json();
return boardJson
}
var board = new Board(); 

export default {
  data() {
    return {
      pieceLocation: board.pieceLocation,
      totalTiles: board.totalTiles,
      selectedTile: board.selectedTile,
      turnText: board.turnText,
      getIndex: board.getIndex.bind({}),
      getPiece: board.getPiece.bind({}),
      selected: board.selected.bind({}),
    }
  },
  mounted() {
      waitForBoard().then(data=> {board.setPieceLoc(data); this.pieceLocation = board.pieceLocation})
  }
}
</script>

