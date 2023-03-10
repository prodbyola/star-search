import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { 
        path: '', 
        component: () => import('pages/IndexPage.vue'), 
        name: 'home',
      },
      {
        path: 'demo',
        component: () => import('pages/demo/MovieList.vue'),
        name: 'movie-list',
        
      },
      { 
        path: 'demo/comments/:movieID', 
        component: () => import('pages/demo/MovieComments.vue'), 
        name: 'movie-comments' 
      },
      { 
        path: 'demo/characters/:movieID', 
        component: () => import('src/pages/demo/MovieCharacters.vue'), 
        name: 'movie-characters' 
      },
      {
        path: 'docs',
        component: () => import('pages/ApiDocs.vue'),
        name: 'docs'
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
