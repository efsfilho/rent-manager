<template>
  <v-dialog
    v-model="visible"
    persistent
    width="200"
  >
    <v-row justify="center">
      <v-col cols="12" xs="12" sm="12" md="12" lg="12" xl="12">
        <v-card>

          <v-card-title class="text-center mx-5 my-2">
            <span>Aguarde</span>
          </v-card-title>

          <v-card-text>
            <v-progress-linear
              color="grey"
              indeterminate
              rounded
              height="5"
              class="mb-3"
            ></v-progress-linear>
          </v-card-text>

        </v-card>
      </v-col>
    </v-row>
  </v-dialog>
</template>

<script lang="ts">

  export default {
    props: {
      timeout: {
        type: Number,
        default: 1000, // millisecs
      }
    },
    data: () => ({
      visible: true,
    }),
    created() {
      if (this.timeout > 0) {
        setTimeout(() => {
          this.visible = false;
        }, this.timeout)
      }
    },
    methods:{
      closeNotification() {
        this.$emit("close");
      }
    }
  }
</script>