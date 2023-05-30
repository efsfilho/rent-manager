<template>
  <v-card>
    <v-card-title>Inquilinos</v-card-title>

    <v-card-text>

      <!-- TOP BUTTONS -->
      <v-card-actions>
        <!-- DEBUG -->
        <v-btn v-if="debugMode" class="mb-2" variant="tonal" @click="getTenants">
          Update
        </v-btn>
        <!-- DEBUG -->

        <v-spacer></v-spacer>
        <v-btn class="mb-2" @click="openTenantForm">
          Adicionar
        </v-btn>
      </v-card-actions>

      <v-expansion-panels>
        <v-expansion-panel
          v-for="tenant in tenantsStore.getTenants"
          :key="tenant.id"
        >
          <v-expansion-panel-title>
            <template v-slot:default="{ expanded }">
              <v-row no-gutters>
                <v-col cols="10" class="d-flex align-center">
                  <span v-if="!expanded" :key="tenant.id">
                    {{ tenant.name }}
                    <!-- {{ 'salkjdfhslfjksajhlkasjflkasjflksajflksjaflksjasldkjaslkdj' }} -->
                  </span>

                </v-col>
                <v-col class="d-flex align-center">
                  <v-btn v-if="debugMode" @click="deleteFromList($event, tenant.id)" size="x-small" class="py-0 my-0 d-flex justify-start" variant="tonal">
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
                    :model-value="tenant.name"
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
                    :model-value="applyMask(tenant.cpf, '###.###.###-##')"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

                <!-- RG -->
                <v-col cols="12" sm="4">
                  <v-text-field
                    label="RG"
                    :model-value="tenant.rg"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

                <!-- DATA NASC -->
                <v-col cols="12" sm="4">
                  <v-text-field
                    label="Data de Nascimento"
                    :model-value="getDate(tenant.birth_date)"
                    variant="outlined"
                    readonly
                  ></v-text-field>
                </v-col>

              </v-row>

              <v-row no-gutters>
                <v-col cols="12" class="d-flex justify-end">
                  <v-btn variant="text" @click="openTenantForm(tenant.id)">
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
        selectedTenant = '';
        showTenantDialog = false;
        getTenants();
      }"
    ></tenant-form>

  </v-card>
</template>

<script lang="ts">
  import { inject } from 'vue';
  import { Utils } from '@/utils/utils';
  import moment from 'moment';
  import TenantForm from '@/components/TenantForm.vue';
  import { Tenant, useTenantStore } from '@/store/tenants';
  
  export default {

    components: {
      TenantForm,
    },
  
    setup(){
      return {
        // injected var/functions provided by layouts/Home.vue
        debugMode: <Boolean> inject('debugMode'),
        showLoading: <Function> inject('showLoading'),
        showNotification: <Function> inject('showNotification'),
        applyMask: Utils.applyMask
      }
    },
  
    data: () => ({
      tenantsStore: useTenantStore(),
      showTenantDialog: false,
      selectedTenant: <string> ''
    }),

    created() {
      this.getTenants();
    },

    methods: {
      async getTenants() {
        await this.showLoading(true);
        try {
          await this.tenantsStore.updateStore();
        } catch (err) {
          if (this.debugMode) {
            console.log(err);
            this.showNotification('debug', err);
          }

          this.showNotification('error', 'Não foi possível atualizar a tela.');
        }
        this.showLoading(false);
      },

      // Used for edit and add
      openTenantForm(tenantId: string) {
        this.showTenantDialog = true;
        this.selectedTenant = tenantId || '';
      },

      deleteFromList(e:any, id:string) {
        e.stopPropagation();
        let tenant = this.tenantsStore.getTenantById(id) as Tenant;
        this.tenantsStore.delete(tenant);
      },

      getDate(uTime: number) {
        return moment(uTime, 'X').format('DD/MM/YYYY');
      },
    },
  }
</script>