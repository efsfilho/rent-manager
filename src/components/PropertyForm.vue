<template>

  <!-- v-model="dialog" -->
  <v-dialog
    persistent
    activator="parent"
  >
    <v-row justify="center">
      <v-col cols="12" xs="12" sm="8" md="6" lg="6" xl="4">
        <v-form ref="formInput">
          <v-card>
            <v-card-title class="ml-5 mt-5">
              <span class="text-h5">{{ propertyId === null ? "Nova" : "Editar"}} Propriedade</span>
            </v-card-title>
            <v-card-text>
              <v-container>

                <!-- NOME -->
                <v-row>
                  <v-col cols="12" sm="12">
                    <v-text-field
                      label="Nome"
                      v-model="property.name"
                      required
                      :rules="[rule]"
                    ></v-text-field>
                  </v-col>
                </v-row>

                <!-- ENDEREÇO -->
                <v-row>
                  <v-col cols="12" sm="12">
                    <v-text-field
                      label="Endereço"
                      v-model="property.address"
                      required
                      :rules="[rule]"
                    ></v-text-field>
                  </v-col>

                </v-row>
              </v-container>
              <small>*indicates required field</small>
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

          <v-dialog
            v-model="loadingDialog"
            :scrim="false"
            persistent
            width="200"
          >
            <v-card
              
            >
              <v-card-title class="mx-5 my-2">
                <span>Aguarde</span>
                <!-- Aguarde -->
              </v-card-title>
              <v-card-text>

                <v-progress-linear
                  color="grey"
                  indeterminate
                  rounded
                  height="5"
                  class="mb-3"
                ></v-progress-linear>
                <!-- <div class="text-center">
                  <v-progress-circular
                    :width="4"
                    :size="70"
                    color="red"
                    indeterminate
                  ></v-progress-circular>
                </div> -->
              </v-card-text>
            </v-card>
          </v-dialog>

        </v-form>
      </v-col>
    </v-row>
  </v-dialog>
  
</template>

<script lang="ts">

  import { Property, usePropertiesStore } from '@/store/properties';
  import moment from 'moment';
  import { inject } from 'vue';
  export default {
    // props: [ "form" ],
    props: [ 'update','propertyId' ],
    setup(){
      const showNotification = <Function> inject('showNotification')
      return {
        showNotification
      }
    },
    data: () => ({
      propertiesStore: usePropertiesStore(),
      loadingDialog: false,
      logout: <any> null,
      property: {
        name: 'Casas TESTE',
        address: 'RUA ABC',
      } as Property,
      // validator: new TenantValidator()
    }),
    computed:{
      isEditing() {        
        return this.propertyId !== undefined && this.propertyId !== null;
      },
    },
    created() {     
      if(this.propertyId !== null) {
        let property = this.propertiesStore
          .getPropertyById(this.propertyId) as Property;
        // console.log("property", property);
        this.property = {...property};
        // this.birthDateInput = moment(tenant.birth_date, 'X').format('DD/MM/YYYY');
      }
    },
    methods: {
      async save() {
        const { valid } = await (this.$refs.formInput as any).validate();
        if (valid) {
          try {
            this.loadingDialog = true;
            
            if(this.isEditing) {
              await this.propertiesStore.update({...this.property});
            } else {
              await this.propertiesStore.add({...this.property});
            }

            this.showNotification('success', 'Propriedade salva com sucesso!')
            this.$emit('close');
          } catch (err) {
            // console.log(err)
            this.loadingDialog = false;
            this.showNotification('error', 'Não foi possível salvar a propriedade.')
          }
        }
      },

      async remove() {
        this.loadingDialog = true;
        try {
          if(this.propertyId !== null) {
            await this.propertiesStore.delete({...this.property});
          }
          this.showNotification('success', 'Propriedade excluída com sucesso!')
          this.$emit('close');
        } catch (err) {
          // console.log(err)
          this.loadingDialog = false;
          this.showNotification('error', 'Não foi possível excluir a propriedade.')
        }
      },
      rule(value: string) {
        if (value)
          return true
        return 'Este campo deve ser preenchido.'
      },
      // allowOnlyNumber(e:any) {
      //   if(/\d/.test(e.key)) {
      //     return true;
      //   } else {
      //     e.preventDefault();
      //   }
      // },

      // formatName() {
      //   this.tenant.name = this.tenant.name.slice(0,50);
      // },

      // formatCpf() {
      //   let cpf = this.cpfInput.slice(0,13);
      //   this.cpfInput = Utils.applyMask(cpf, '###.###.###-##')
      //   // this.tenant.cpf = Utils.applyMask(cpf, '##.##)##-#$@##')
      // },

      // formatRg() {
      //   this.tenant.rg = this.tenant.rg.slice(0,12);
      // },

      // formatDate() {
      //   let date = this.birthDateInput.slice(0,9);
      //   this.birthDateInput = Utils.applyMask(date, '##/##/####');
      // },
      
      // nameRule(name: any) {
      //   return this.validator.isNameValid(name) || 'Nome inválido';
      // },

      // cpfRule(cpf: any) {
      //   return this.validator.isCpfValid(cpf) || 'CPF inválido';
      // },

      // rgRule(rg: any) {
      //   return this.validator.isRgValid(rg) || 'RG inválido';
      // },

      // dateRule(date: string) {
      //   return this.validator.isBirthDateValid(date) || 'Data inválida';
      // }
    }
  }

</script>