<template>

  <v-dialog
    persistent
    activator="parent"
  >
    <v-row justify="center">
      <v-col cols="12" xs="12" sm="8" md="6" lg="6" xl="4">
        <v-form ref="formInput">
          <v-card>

            <v-card-title class="ml-5 mt-5">
              <span class="text-h5">{{isEditing ? "Editar" : "Novo"}} Inquilino</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>

                  <!-- NOME -->
                  <v-col cols="12">
                    <v-text-field
                      label="Nome"
                      v-model="tenant.name"
                      required
                      @keydown="formatName"
                      :rules="[nameRule]"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row>

                  <!-- CPF -->
                  <v-col cols="12" sm="4">
                    <v-text-field
                      label="CPF"
                      required
                      v-model="cpfInput"
                      @keydown="formatCpf"
                      @keypress="allowOnlyNumber"
                      :rules="[cpfRule]"
                    ></v-text-field>
                  </v-col>

                  <!-- RG -->
                  <v-col cols="12" sm="4">
                    <v-text-field
                    label="RG"
                    required
                    v-model="tenant.rg"
                    @keydown="formatRg"
                    :rules="[rgRule]"
                    ></v-text-field>
                  </v-col>

                  <!-- DATA NASC -->
                  <v-col cols="12" sm="4">
                    <v-text-field
                      label="Data de Nascimento"
                      required
                      v-model="birthDateInput"
                      @keydown="formatDate"
                      @keypress="allowOnlyNumber"
                      :rules="[dateRule]"
                    ></v-text-field>
                  </v-col>

                </v-row>
              </v-container>
              <small>*indica campos obrigatórios</small>
            </v-card-text>

            <v-card-actions>

              <!-- EXCLUIR -->
              <v-btn
                v-if="isEditing"
                color="red-darken-1"
                variant="text"
                @click="remove"
              >
                Excluir
              </v-btn>

              <v-spacer></v-spacer>

              <!-- CANCELAR -->
              <v-btn
                color="blue-darken-1"
                variant="text"
                @click="$emit('close');"
              >
                Cancelar
              </v-btn>

              <!-- SALVAR -->
              <v-btn
                color="blue-darken-1"
                variant="text"
                @click="save"
              >
                Salvar
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-form>
      </v-col>
    </v-row>
  </v-dialog>
</template>

<script lang="ts">
  import { Utils } from '@/utils/utils';
  import { inject } from 'vue';
  import { TenantValidator } from '@/modules/Tenant';
  import { Tenant, useTenantStore } from '@/store/tenants';
  import moment from 'moment';
  
  export default {

    props: {
      tenantId: {
        type: String,
        default: '',
      }
    },

    setup(){
      return {
        // injected var/functions provided by layouts/Home.vue
        debugMode: <Boolean> inject('debugMode'),
        showLoading: <Function> inject('showLoading'),
        showNotification: <Function> inject('showNotification'),
      }
    },

    data: () => ({
      tenantStore: useTenantStore(),
      tenant: <Tenant>{},
      cpfInput: '',
      birthDateInput: '',
      validator: new TenantValidator()
    }),

    created() {
      if(this.isEditing) {
        let tenant = this.tenantStore.getTenantById(this.tenantId) as Tenant;
        this.tenant = structuredClone(tenant);
        this.cpfInput = this.tenant.cpf;
        this.formatCpf();        
        this.birthDateInput = moment(tenant.birth_date, 'X').format('DD/MM/YYYY');
      } else {

        // DEBUG
        if (this.debugMode) {
          this.tenant.name = 'TESTE';
          this.tenant.rg = '12345611';
          this.tenant.birth_date = 0;
          this.cpfInput = '092.421.600-08'
          this.birthDateInput = '12/12/1991';
        }
      }
    },

    computed: {
      isEditing() {        
        return Utils.isValidObjectId(this.tenantId);
      },
    },

    watch: {
      cpfInput() {
        this.tenant.cpf = this.cpfInput.replace(/\.|\-/g, "");
      },

      // convert DD/MM/YYYY format to timestamp seconds
      birthDateInput(val: string) {
        let date = moment(val, 'DD/MM/YYYY', true);

        if (date.isValid()) {
          this.tenant.birth_date = date.unix();
        } else {
          this.tenant.birth_date = 0;
        }
      }
    },

    methods: {
      async save() {
        const { valid } = await (this.$refs.formInput as any).validate();

        if (valid) {
          this.showLoading(true);

          try {
            let tenant = structuredClone(this.tenant);

            if(this.isEditing) {
              await this.tenantStore.update(tenant);
            } else {
              await this.tenantStore.add(tenant);
            }
            this.showNotification('success', 'Inquilino salvo com sucesso!');
            this.$emit('close');
          } catch (err) {
            if (this.debugMode) {
              console.log(err);
              this.showNotification('debug', err);
            }

          this.showLoading(false);
          this.showNotification('error', 'Não foi possível salvar o inquilino.');
          }

          this.showLoading(false);
        }
      },

      async remove() {
        this.showLoading(true);

        try {

          if(this.isEditing) {
            await this.tenantStore.delete({...this.tenant});
            this.showNotification('success', 'Inquilino excluído com sucesso!');
            this.$emit('close');
          }

        } catch (err) {

          if (this.debugMode) {
            console.log(err);
            this.showNotification('debug', err);
          }

          this.showNotification('error', 'Não foi possível excluir o inquilino.');
        }

        this.showLoading(false);
      },

      

      allowOnlyNumber(e:any) {
        if(/\d/.test(e.key)) {
          return true;
        } else {
          e.preventDefault();
        }
      },

      formatName() {
        this.tenant.name = this.tenant.name.slice(0,50);
      },

      formatCpf() {
        let cpf = this.cpfInput.slice(0,13);
        this.cpfInput = Utils.applyMask(cpf, '###.###.###-##');
      },

      formatRg() {
        this.tenant.rg = this.tenant.rg.slice(0,12);
      },

      formatDate() {
        let date = this.birthDateInput.slice(0,9);
        this.birthDateInput = Utils.applyMask(date, '##/##/####');
      },
      
      nameRule(name: any) {
        return this.validator.isNameValid(name) || 'Nome inválido';
      },

      cpfRule(cpf: any) {
        return this.validator.isCpfValid(cpf) || 'CPF inválido';
      },

      rgRule(rg: any) {
        return this.validator.isRgValid(rg) || 'RG inválido';
      },

      dateRule(date: string) {
        return this.validator.isBirthDateValid(date) || 'Data inválida';
      }
    }
  }

</script>