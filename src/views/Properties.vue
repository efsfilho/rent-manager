<template>
  <v-card>
    <v-card-title>Propriedades</v-card-title>

    <v-card-text>
      <v-card-actions>
        <v-btn v-if="debugMode" class="mb-2" variant="tonal" @click="getProperties">
          Update
        </v-btn>
        <v-spacer></v-spacer>
        <v-btn class="mb-2" @click="addProperty">
          Adicionar
        </v-btn>
      </v-card-actions>

    
      <v-expansion-panels>
        <v-expansion-panel
          v-for="t in propertiesStore.properties"
          :key="t.id"
        >
          <v-expansion-panel-title>
            <template v-slot:default="{ expanded }">
              <v-row no-gutters>
                <v-col cols="10" class="d-flex justify-start">
                  <span v-if="!expanded" :key="t.id">
                    {{ t.name }}
                  </span>
                </v-col>
                <v-col class="py-0 my-0">
                  <v-btn
                    v-if="debugMode"
                    @click="deleteFromList($event, t.id)"
                    size="x-small"
                    class="py-0 my-0 d-flex justify-start"
                    variant="tonal"
                  >
                    delete
                  </v-btn>
                </v-col>
              </v-row>
            </template>
          </v-expansion-panel-title>

          <v-expansion-panel-text>
            <v-container>

              <!-- NOME -->
              <v-row>
                <v-col cols="12" sm="12">
                  <v-text-field
                    label="Nome"
                    :model-value="t.name"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>
              </v-row>

              <!-- ENDEREÇO -->
              <v-row>
                <v-col cols="12" sm="12">
                  <v-text-field
                    label="Endereço"
                    :model-value="t.address"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

                </v-row>

                <v-row no-gutters>
                <v-col cols="12" class="d-flex justify-end">
                  <v-btn variant="text" @click="editProperty(t.id)">
                    Editar
                  </v-btn>
                </v-col>
                </v-row>
            </v-container>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card-text>

    <property-form
      v-if="showPropertyDialod"
      :property-id="selectedProperty"
      @close="() => {
        selectedProperty = null;
        showPropertyDialod = false;
      }"
      @edit-tenant="editProperty"
    >
    </property-form>

  </v-card>
</template>

<script lang="ts">
  import { inject } from 'vue';
  import PropertyForm from '@/components/PropertyForm.vue';
  import { Property, usePropertiesStore } from '@/store/properties';
  
  type Selected = number | null;
  
  export default {
    components: {
      PropertyForm,
    },
    setup(){
      // injected var/functions provided by layouts/Home.vue
      const debugMode = <Boolean> inject('debugMode');
      const showNotification = <Function> inject('showNotification')
      return {
        debugMode,
        showNotification
      }
    },
    data: () => ({
      propertiesStore: usePropertiesStore(),
      showPropertyDialod: false,
      selectedProperty: <Selected> null,
    }),
    created() {
      this.getProperties();
    },
    methods: {
      getProperties() {
        try {
          this.propertiesStore.get();
        } catch (err) {
          this.showNotification('error', 'Não foi possível atualizar a tela.')
        }
      },
      addProperty() {
        this.showPropertyDialod = true;
      },
      deleteFromList(e:any, id:number) {
        e.stopPropagation()        
        let property = this.propertiesStore.getPropertyById(id) as Property;
        this.propertiesStore.delete(property);
      },

      async editProperty(propertyId: number) {
        this.showPropertyDialod = true;
        this.selectedProperty = propertyId;
      },
    },
  }
</script>