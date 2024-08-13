import { reactive } from 'vue'

// export const store = reactive({
//   count: 0
// });
// const route = useRoute();
import { useRouter, useRoute } from 'vue-router';
const routerPush = (routerName) => {
  const router = useRouter();
  router.push({ name: routerName });
}
const test = reactive({
  count: 0,
  title: '',
  increment() {
    this.count++
  },
  setNewBlockMenu() {
    this.title = 'New Block'
  },
  goMain() {
    // router.push({ name:'home' })
    routerPush('home');
  },
  newBlock() {
    console.log('newblock')
    // router.push({ name:'new-block' })
    routerPush('new-block')
  },
  teste1() {
    const router = useRouter();
    router.push({ name:'teste1' })
    routerPush('teste1')
  },
  test2() {
    // router.push({ name:'teste2' })
    // debugger
    routerPush('teste2')
  }
});

const menu = {
  goMain() {
    router.push({ name:'home' })
    // routerPush('home');
  },
  newBlock() {
    console.log('newblock')
    // router.push({ name:'new-block' })
    routerPush('new-block')
  },
  teste1() {
    // const router = useRouter();
    // router.push({ name:'teste1' })
    routerPush('teste1')
  },
  test2() {
    // router.push({ name:'teste2' })
    routerPush('teste2')
  }
}

export { test, menu }
// command: () => router.push({ name:'new-block', props: { title: 123 } })
//     // command: () => router.push({ name:'new-block', params: { title: 123 } })
//     // command: () => console.log('asdasdsdsad')
// },
// {
//     label: 'Teste 1',
//     icon: 'pi pi-home',
//     command: () => router.push({ name:'teste1' })
// },
// {
//     label: 'Teste 2',
//     icon: 'pi pi-star',
//     command: () => router.push({ name:'teste2' })