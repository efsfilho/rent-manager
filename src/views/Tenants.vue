<template>
  <v-card>
    <v-card-title>Tenant</v-card-title>

    <v-card-text>
      <v-card-actions>
        <v-btn class="mb-2" variant="tonal" @click="teste();">
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
                <v-col cols="10" class="d-flex justify-start">
                  <span v-if="expanded" :key="t.id">
                    Full {{ t.name }}
                  </span>
                  <span v-else key="1">
                    {{ t.name }}
                  </span>
                </v-col>
                <v-col class="py-0 my-0">
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
  // import { Operation } from '@/enum';
  import { Utils } from '@/utils/utils';
  import moment from 'moment';
  import TenantForm from '@/components/TenantForm.vue';
  import { Tenant, useTenantStore } from '@/store/tenants';
  type Selected = number | null;

  export default {
    components: {
      TenantForm,
    },
    data: () => ({
      tenantStore: useTenantStore(),
      showTenantDialog: false,
      selectedTenant: <Selected> null
    }),
    created() {
      this.tenantStore.get();
      // this.tenantStore.add({ id: 0, name: 'TESTE 111', cpf: '48778634865', rg: '11111111', birth_date: 1445738400 });
      // this.tenantStore.add({ id: 1, name: 'TESTE 222', cpf: '07893646039', rg: '22222222', birth_date: 1445732326 });
      // this.tenantStore.add({ id: 2, name: 'TESTE 333', cpf: '26294154227', rg: '33333333', birth_date: 1445732326 });
    },
    methods: {
      deleteFromList(e:any, id:number) {    
        e.stopPropagation()        
        let tenant = this.tenantStore.getTenantById(id) as Tenant;
        this.tenantStore.delete(tenant);
      },
      
      teste() {
        this.tenantStore.get()
      },
      addTenant() {
        this.showTenantDialog = true;
      },

      async editTenant(tenantId: number) {
        this.showTenantDialog = true;
        this.selectedTenant = tenantId;
      },

      getDate(uTime: number) {
        return moment(uTime, 'X').format('DD/MM/YYYY');
      },
      applyMask(value: string, mask: string) {
        return Utils.applyMask(value, mask);
      },

    },
    watch: {
      // form() {
      //   console.log('form: ', this.showTenantDialog);
      // }
    }
  }
</script>