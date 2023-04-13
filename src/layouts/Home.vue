<template>
  <!-- <v-card> -->
    <v-layout>
      <!-- <v-card class="mx-auto" variant="outlined"> -->
      <!-- <v-app-bar color="surface-variant" title="Application bar"></v-app-bar> -->
      <!-- <Menu></Menu> -->
      <right-menu></right-menu>

      <!-- <v-card class="mx-auto" max-width="344" variant="outlined"> -->
      <v-main style="min-height: 400px;" >
        <v-container >
          <!-- <v-row  >
            <v-col> -->

              <router-view ></router-view>
            <!-- </v-col>
          </v-row> -->
          
        </v-container>
      </v-main>
    <!-- </v-card> -->
  <!-- </v-card> -->
    </v-layout>
  <!-- <v-app>
    <Menu></Menu>

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col>
            <Tenant></Tenant>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app> -->
</template>

<script lang="ts" >
  import RightMenu from '@/components/RightMenu.vue'
  
  export default {
    data: () => ({

      items: [
        { title: 'Home', icon: 'mdi-home-city' },
        { title: 'My Account', icon: 'mdi-account' },
        { title: 'Users', icon: 'mdi-account-group-outline' },
      ],
      rail: false,
    }),
    components: { RightMenu }
  }
</script>
