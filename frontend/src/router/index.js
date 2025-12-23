/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import { routes } from 'vue-router/auto-routes'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

// دالة للتحقق من تسجيل الدخول
function isAuthenticated() {
  return localStorage.getItem('isAuthenticated') === 'true'
}

// Navigation Guard - حماية الصفحات
router.beforeEach((to, from, next) => {
  // الصفحات المسموح الوصول إليها بدون تسجيل دخول
  const publicPages = ['/login']
  const isPublicPage = publicPages.includes(to.path)
  
  // إذا كان المستخدم غير مسجل دخول ومحاولة الوصول لصفحة محمية
  if (!isAuthenticated() && !isPublicPage) {
    // إعادة التوجيه إلى صفحة تسجيل الدخول
    next('/login')
  } 
  // إذا كان المستخدم مسجل دخول ومحاولة الوصول لصفحة تسجيل الدخول
  else if (isAuthenticated() && isPublicPage) {
    // إعادة التوجيه إلى الصفحة الرئيسية
    next('/')
  } 
  // السماح بالوصول
  else {
    next()
  }
})

// Workaround for https://github.com/vitejs/vite/issues/11804
router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (localStorage.getItem('vuetify:dynamic-reload')) {
      console.error('Dynamic import error, reloading page did not fix it', err)
    } else {
      console.log('Reloading page to fix dynamic import error')
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router
