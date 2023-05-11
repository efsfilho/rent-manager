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
  
    <!-- v-model="visible" -->
    <v-snackbar
      v-model="visible"
      :color="notificationType"
      variant="elevated"
      location="top"
      max-height="100%"
      :vertical="messageLines.length > 1"
    >
      <div v-if="messageLines.length > 1">
        <div v-for="(line, i) in messageLines">
          <div v-if="i == 0" class="text-subtitle-1">{{i+''+line}}</div>
          <p v-else>{{i+''+line}}</p>
        </div>
      </div>
      <div v-else>
        <div class="text-subtitle-1">{{messageLines[0]}}</div>
      </div>

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
      message: Array<String>,
      notificationType: String,
    },

    data: () => ({
      visible: true,
    }),

    computed: {
      messageLines():Array<String> {
        if (Array.isArray(this.message)) {
          return this.message;
        } else {
          return [this.message || ''];
        }
      }
    },

    methods:{
      closeNotification() {
        this.$emit("close");
      }
    }
  }
</script>