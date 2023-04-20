<template>
  <div></div>
  <v-navigation-drawer
        v-model="drawer"
        :rail="rail"
        :width="250"
        permanent
        @click="rail = false"
      >
        <v-list-item
          prepend-avatar="https://randomuser.me/api/portraits/men/85.jpg"
          title="John Leider"
          nav
        >
          <template v-slot:append>
            <v-btn
              variant="text"
              icon="mdi-chevron-left"
              @click.stop="rail = !rail"
            ></v-btn>
          </template>
        </v-list-item>

        <v-divider></v-divider>

        <v-list density="compact" nav>
          <v-list-item
            v-for="link in links"
            :key="link.name"
            :title="link.name"
            :value="link.name"
            :prepend-icon="link.icon"
            :to="link.link"
          >
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
</template>

<script lang="ts">
  export default {
    props: ['menuLinks'],
    data: () => ({
      drawer: true,
      // links: [
      //   // { name: "Alugueis", link: "/rents", icon:"mdi-currency-usd"},
      //   // { name: "Inquilinos", link: "/tenants", icon:"mdi-account-box" },
      //   // { name: "Imóveis", link: "/properties", icon:"mdi-home-city" },
      // ],
      rail: false,
    }),
    computed: {
      links() {
        return this.menuLinks
      }
    }
  }
</script>
