<template>
  <v-card>
    <v-card-title>{{title}}</v-card-title>
    <!-- <v-card-actions>
      <v-btn color="red-darken-1">Excluir</v-btn>
      <v-spacer></v-spacer>
      <v-btn color="blue-darken-1">Cancelar</v-btn>
      <v-btn color="blue-darken-1">Salvar</v-btn>
    </v-card-actions> -->
    <v-card-text>
      <slot></slot>
    </v-card-text>
    
  </v-card>
</template>

<script lang="ts">

  export default {
    // props: {
    //   title: String
    // },
  props:["title", "cont"],
    data: () => ({
    
    })
  }
</script>