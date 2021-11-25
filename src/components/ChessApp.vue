<template>
  <div class="boardBorder">
    <table class="board">
      <tr v-for="row in totalTiles" class="row">
          <td v-for="column in totalTiles" v-bind:id="'tile'+getIndex(row, column)" @click="selectedTile=selected(row,column,pieceLocation,selectedTile)" class= "column">
            <div :class="{'active': getIndex(row, column) == selectedTile}">
                <img :src="getPiece(row,column,pieceLocation)" v-bind:id="'img'+getIndex(row, column)">
            </div>
          </td>
      </tr>
    </table>
  </div>
</template>

<script>
import Board from "../classes/Board.js";
var board = new Board()

export default {
  data() {
    return {
      pieceLocation: board.pieceLocation,
      totalTiles: board.totalTiles,
      selectedTile: board.selectedTile,
      getIndex: board.getIndex.bind({}),
      getPiece: board.getPiece.bind({}),
      selected: board.selected.bind({}),
    }
  }
}
</script>

