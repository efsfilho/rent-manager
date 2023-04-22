<template>
  <v-card>
    <v-card-title>Alugueis</v-card-title>

    <v-card-text>
      <v-card-actions>
        <!-- <v-btn class="mb-2" variant="tonal" @click="tenantStore.get();"> -->
        <v-btn class="mb-2" variant="tonal">
          Update
        </v-btn>
        <v-spacer></v-spacer>
        <!-- <v-btn class="mb-2" @click="addTenant"> -->
        <v-btn class="mb-2" >
          Adicionar
        </v-btn>
      </v-card-actions>

    
      <v-expansion-panels>
        <v-expansion-panel
          v-for="t in items"
          :key="t.id"
        >
        <!-- <v-expansion-panel
          v-for="t in tenants"
          :key="t.id"
        > -->
          <!-- text="Lorem ipsum dolor dsit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat." -->
          <v-expansion-panel-title>
            <template v-slot:default="{ expanded }">
              <v-row no-gutters>
                <v-col cols="10" class="d-flex justify-start">
                  <span v-if="!expanded" :key="t.id">
                    {{ t.name }}
                  </span>
                </v-col>
                <v-col class="py-0 my-0">
                  <!-- <v-btn @click="deleteFromList($event, t.id)" size="x-small" class="py-0 my-0 d-flex justify-start" variant="tonal"> -->
                  <v-btn size="x-small" class="py-0 my-0 d-flex justify-start" variant="tonal">
                    delete
                  </v-btn>
                </v-col>
              </v-row>
            </template>
          </v-expansion-panel-title>

          <v-expansion-panel-text>
            <v-container>
              <v-row>

                <!-- NOME -->
                <v-col cols="12">
                  <v-text-field
                    label="Nome"
                    :model-value="t.name"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

              </v-row>
              <v-row no-gutters>
                <v-col cols="12" class="d-flex justify-end">
                  <!-- <v-btn variant="text" @click="editTenant(t.id)"> -->
                  <v-btn variant="text" >
                    Editar
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card-text>
  </v-card>  

</template>

<script lang="ts">
  // import Card from "@/components/Card.vue"
  export default {
    data: () => ({
      items: [
        { id: 0, name : 'ASDAS'},
        { id: 1, name : 'FDSFSE'},
        { id: 2, name : 'GFHFGH'},
        { id: 3, name : 'GHJFGJHG'},
        { id: 4, name : '&¨$%&%¨'},
      ],
      reveal: false
    }),
    // components: { Card }

  }
</script>