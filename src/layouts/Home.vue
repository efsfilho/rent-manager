<template>
  <v-layout>
      <!-- <v-card class="mx-auto" variant="outlined"> -->
      <!-- <v-app-bar color="surface-variant" title="Application bar"></v-app-bar> -->
      <!-- <Menu></Menu> -->
      <right-menu
        :menu-links="menuLinks"
      ></right-menu>

      <!-- <v-card class="mx-auto" max-width="344" variant="outlined"> -->
    <v-main style="min-height: 400px;" >
      <v-container >
        <!-- <v-row>
          <v-col>
            <v-btn @click="showNotification">Open Snackbar</v-btn>
            <v-btn @click="notificationVisible = false">Close Snackbar</v-btn>
          </v-col>
          <v-col> -->
        <router-view ></router-view>
          <!-- </v-col>
        </v-row> -->

        <!-- NOTIFICATION -->
        <notification
          v-if="notificationVisible"
          :message="notificationMessage"
          :notification-type="notificationType"
          @close="closeNotification"
        >
        </notification>

        <!-- DEBUG -->
        <v-dialog v-model="debugDialogVisible" width="90%">
          <v-card>
            <v-card-title class="text-center">
              <span class="text-h5">DEBUG</span>
            </v-card-title>
            <v-card-text>
              <div v-for="(line, i) in debugMessage">
                <p>{{line}}</p>
              </div>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn @click="debugDialogVisible = false">Close</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
  
      </v-container>
    </v-main>

  </v-layout>
</template>

<script lang="ts" >
  import Notification from '@/components/Notification.vue';
  import RightMenu from '@/components/RightMenu.vue'

  type messageType = 'success' | 'info' | 'warning' | 'error' | 'debug';

  export default {
    data: () => ({
      notificationVisible: false,
      notificationType: '',
      notificationMessage: <string[]> [],
      debugDialogVisible: false,
      debugMessage: <string[]> [],
      count: 0,
      menuLinks: [
        { name: "Alugueis", link: "/rents", icon:"mdi-currency-usd"},
        { name: "Inquilinos", link: "/tenants", icon:"mdi-account-box" },
        { name: "Imóveis", link: "/properties", icon:"mdi-home-city" },
      ],
      rail: false,
    }),
    provide() {
      return {
        debugMode: false,
        showNotification: this.showNotification
      }
    },
    methods: {
      showNotification(type: messageType, message: string | string[] | Error) {
        let messageContent = [];
        if (this.isError(message)) {
          messageContent = this.getTextFromError(message);
        } else {
          if (!Array.isArray(message)) {
            message = [message];
          }
          messageContent = message;
        }

        if (type === 'debug') {
          this.debugDialogVisible = true;
          this.debugMessage = messageContent;
        } else {

          this.notificationType = type;
          if (this.notificationVisible) {
            this.notificationVisible = false;
          }
          setTimeout(() => this.notificationVisible = true, 300);
          this.notificationMessage = messageContent;
        }
      },

      isError(obj: any):boolean {
        let hasErrorProperties = obj.hasOwnProperty('stack') && obj.hasOwnProperty('message');
        return hasErrorProperties || obj instanceof Error
      },

      getTextFromError(error: any) {
        let txt = [];
        let expectedProps = ['name', 'code', 'stack']
        for (const prop of expectedProps) {
          if (expectedProps.includes(prop))
            txt.push(`${prop}: ${error[prop]}\n`);
        }
        return txt;
      },

      closeNotification() {
        this.notificationVisible = false;
      }
    },
    watch: {
      notificationVisible() {
        this.count++;
      }
    },
    components: { 
      Notification,
      RightMenu,
    }
  }
</script>
