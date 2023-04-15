<template>
  <v-layout>
      <!-- <v-card class="mx-auto" variant="outlined"> -->
      <!-- <v-app-bar color="surface-variant" title="Application bar"></v-app-bar> -->
      <!-- <Menu></Menu> -->
      <right-menu></right-menu>

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
        
        <notification
          v-if="notificationVisible"
          :message="notificationMessage"
          :type-color="notificationType"
          @close="closeNotification"
        >
          <!-- type-color="notificationType" -->
      </notification>
      </v-container>
    </v-main>

  </v-layout>
</template>

<script lang="ts" >
  import Notification from '@/components/Notification.vue';
  import RightMenu from '@/components/RightMenu.vue'

  type messageType = 'success' | 'info' | 'warning' | 'error';

  export default {
    data: () => ({
      notificationVisible: false,
      notificationType: '',
      notificationMessage: '',
      count: 0,
      items: [
        { title: 'Home', icon: 'mdi-home-city' },
        { title: 'My Account', icon: 'mdi-account' },
        { title: 'Users', icon: 'mdi-account-group-outline' },
      ],
      rail: false,
    }),
    provide() {
      return {
        showNotification: this.showNotification
      }
    },
    methods: {
      showNotification(type: messageType, message: string) {
        if (this.notificationVisible) {
          this.notificationVisible = false;
        }

        setTimeout(() => {
          this.notificationVisible = true;
        }, 300);
        
        this.notificationType = type;
        this.notificationMessage = message
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
