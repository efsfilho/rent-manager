<template>
  <!-- <v-card>
    <v-card-title>{{title}}</v-card-title>
    <v-card-actions>
      <v-btn color="red-darken-1">Excluir</v-btn>
      <v-spacer></v-spacer>
      <v-btn color="blue-darken-1">Cancelar</v-btn>
      <v-btn color="blue-darken-1">Salvar</v-btn>
    </v-card-actions>
    <v-card-text>
      <slot></slot>
    </v-card-text>
  </v-card> -->


  <div class=" ">

    <v-snackbar
      v-model="visible"
      :color="typeColor"
      variant="elevated"
      location="top"
      >
      <!-- <div class="text-caption">Erro</div>  -->
      <div class="text-subtitle-1">
        <!-- <v-icon
          start
          icon="error"
        ></v-icon> -->
        {{ message }}
      </div>
      <!-- <p class="text-body-1">This is a snackbar measdasddsdsassage</p> 
      <div class="text-subtitle-2">This is a snackbar message</div>  -->
      <!-- <v-row>
        <v-col cols="2">
          <div class="text-body-1">Erro</div>
          <div class="text-body-2">Erro</div>
        </v-col>
        <v-col cols="10">
          
          <div class="text-body-2">This is a snackbar message</div>
        </v-col>
      </v-row> -->
      <!-- {{ message }} -->
      <template v-slot:actions>
        <v-btn
          variant="text"
          @click="closeNotification"
        >
          Fechar
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script lang="ts">

  export default {
    props: {
      message: String,
      typeColor: String,
    },
    data: () => ({
      visible: true,
    }),
    methods:{
      closeNotification() {
        this.$emit("close");
      }
    }
  }
</script>