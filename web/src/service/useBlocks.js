// import { useQueryClient, useQuery, useMutation } from '@tanstack/vue-query';
// // // Acess QueryClient instance
// const queryClient =  useQueryClient();

// const getTodos = async() => {
//     // const p = await axios.get('https://jsonplaceholder.typicode.com/posts');
//     const p = await axios.get('http://localhost:3000/blocks');
//     console.log('ppp>>>>>', JSON.stringify(p.data))
//     return p.data;
//     // return teste
// }
// // Query
// const { isLoading, isError, data, error } = useQuery({
//   queryKey: ['todos'],
//   queryFn: getTodos,
// });

// // const postTodo = async(e) => {
// //     console.log('postTodo', e);
// //     teste.push(e)
// // }
// // // Mutation
// // const mutation = useMutation({
// //   mutationFn: postTodo,
// //   onSuccess: () => {
// //     // Invalidate and refetch
// //     // console.log('DDASDAS');
// //     value.value = '';
// //     queryClient.invalidateQueries({ queryKey: ['todos'] });
// //   },
// // });

// // const { isLoading, isError, data, error } = useQuery({
// const { isPending, isSuccess, mutate } = useMutation({
//   mutationFn: (newTodo) => axios.post('http://localhost:3000/blocks', newTodo),
//     onSuccess: () => {
//     // Invalidate and refetch
//     // console.log('DDASDAS');
//     value.value = '';
//     queryClient.invalidateQueries({ queryKey: ['blocks'] });
//   },
// });

// function onButtonClick() {
//     mutate({ id: new Date(), title: 'Do Laundry' })
// }

export const useBlocks = () => {


}

import { useQuery } from '@tanstack/vue-query';
import { ref } from 'vue'

// const fetchData = async (filter) => {
//   let url = `https://jsonplaceholder.typicode.com/todos/${filter ? filter : ''}`
//   try {
//     const res = await fetch(url);
//     const data = await res.json();
//     return Array.isArray(data)? data : [data];
//   } catch (ex) {
//     console.log(ex);
//   }
// };

const getTodos = async() => {
    const p = await axios.get('http://localhost:3000/blocks');
    console.log('ppp>>>>>', JSON.stringify(p.data))
    return p.data;
}

export const useTodos = (enabled = true) => {
//   const filtersref = ref();
//   const updatefilters =(filters) => {
//     filtersref.value = filters;
//     refetch();
//   }

  const { data, isLoading, isError, error, refetch } = useQuery(
    ['blocks'],
    () => getTodos(), //fetchData(filtersref.value),
    { enabled, }
  );

  return { data, isLoading, isError, error, updatefilters };
};
