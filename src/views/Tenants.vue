<template>
  <v-card>
    <v-card-title>Inquilinos</v-card-title>

    <v-card-text>
      <v-card-actions>
        <v-btn class="mb-2" variant="tonal" @click="tenantStore.get();">
          Update
        </v-btn>
        <v-spacer></v-spacer>
        <v-btn class="mb-2" @click="addTenant">
          Adicionar
        </v-btn>
      </v-card-actions>

    
      <v-expansion-panels>
        <v-expansion-panel
          v-for="t in tenantStore.getTenants"
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
                <v-col cols="10" class="d-flex align-center">
                  <span v-if="!expanded" :key="t.id">
                    {{ t.name }}
                    <!-- {{ 'salkjdfhslfjksajhlkasjflkasjflksajflksjaflksjasldkjaslkdj' }} -->
                  </span>
                <!-- </v-col>
                <v-col cols="4" class="d-flex align-center"> -->
                  <!-- v-if="t.property_id != 0" -->
                  <v-chip
                    v-if="t.property"
                    class="mx-4"
                    size="small"
                    color="primary"
                    label
                  >
                    <v-icon start icon="mdi-home-city-outline"></v-icon>
                    {{ t.property.name }}
                  </v-chip>
                </v-col>
                <v-col class="d-flex align-center">
                  <v-btn @click="deleteFromList($event, t.id)" size="x-small" class="py-0 my-0 d-flex justify-start" variant="tonal">
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
              <v-row>

                <!-- CPF -->
                <v-col cols="12" sm="4">
                  <v-text-field
                    label="CPF"
                    :model-value="applyMask(t.cpf, '###.###.###-##')"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

                <!-- RG -->
                <v-col cols="12" sm="4">
                  <v-text-field
                    label="RG"
                    :model-value="t.rg"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

                <!-- DATA NASC -->
                <v-col cols="12" sm="4">
                  <v-text-field
                    label="Data de Nascimento"
                    :model-value="getDate(t.birth_date)"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

              </v-row>

              <v-row no-gutters>
                <v-col cols="12" class="d-flex justify-end">
                  <v-btn variant="text" @click="editTenant(t.id)">
                    Editar
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card-text>

    <tenant-form
      v-if="showTenantDialog"
      :tenant-id="selectedTenant"
      @close="() => {
        selectedTenant = null;
        showTenantDialog = false;
      }"
      @edit-tenant="editTenant"
    ></tenant-form>

  </v-card>
</template>

<script lang="ts">
  import { inject } from 'vue';
  import { Utils } from '@/utils/utils';
  import moment from 'moment';
  import TenantForm from '@/components/TenantForm.vue';
  import { Tenant, useTenantStore } from '@/store/tenants';

  type Selected = number | null;
  
  export default {
    components: {
      TenantForm,
    },
    setup(){
      const showNotification = <Function> inject('showNotification');
      const applyMask = Utils.applyMask;
      const showShortName = Utils.showShortName;
      return {
        showNotification,
        applyMask,
        showShortName
      }
    },
    data: () => ({
      tenantStore: useTenantStore(),
      showTenantDialog: false,
      selectedTenant: <Selected> null
    }),
    created() {
      this.getTenants();
    },
    methods: {
      async getTenants() {
        try {
          await this.tenantStore.get();
        } catch (err) {
          console.log(err)
          this.showNotification('error', 'Não foi possível atualizar a tela.')
        }
      },
      addTenant() {
        this.showTenantDialog = true;
      },
      deleteFromList(e:any, id:number) {    
        e.stopPropagation()        
        let tenant = this.tenantStore.getTenantById(id) as Tenant;
        this.tenantStore.delete(tenant);
      },
      async editTenant(tenantId: number) {
        this.showTenantDialog = true;
        this.selectedTenant = tenantId;
      },

      getDate(uTime: number) {
        return moment(uTime, 'X').format('DD/MM/YYYY');
      },
      // applyMask(value: string, mask: string) {
      //   return Utils.applyMask(value, mask);
      // },

    },
    watch: {
      // form() {
      //   console.log('form: ', this.showTenantDialog);
      // }
    }
  }
</script>