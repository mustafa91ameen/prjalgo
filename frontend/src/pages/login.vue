<template>
  <div class="login-page">
    <div class="system-title-wrapper">
      <h1 class="system-title">نظام إدارة المشاريع</h1>
    </div>
    <div class="box">
      <div class="login">
        <div class="loginbx">
          <h2>
            <i class="mdi mdi-lock"></i> تسجيل الدخول
          </h2>
          <form @submit.prevent="handleLogin">
            <div v-if="errorMessage" class="error-message">
              {{ errorMessage }}
            </div>
            <input
              v-model="formData.username"
              type="text"
              placeholder="اسم المستخدم"
              required
              :disabled="loading"
            />
            <input
              v-model="formData.password"
              type="password"
              placeholder="كلمة المرور"
              required
              :disabled="loading"
            />
            <div class="group">
              <a href="#" @click.prevent="handleForgotPassword">نسيت كلمة المرور؟</a>
              <a href="#" @click.prevent="handleRegister">إنشاء حساب جديد</a>
            </div>
            <input 
              type="submit" 
              :value="loading ? 'جاري تسجيل الدخول...' : 'دخول'"
              :disabled="loading"
            />
          </form>
          
          <!-- معلومات بيانات تسجيل الدخول (للتطوير فقط) -->
          <div v-if="showCredentials" class="credentials-info">
            <h3>بيانات تسجيل الدخول المتاحة:</h3>
            <div class="credentials-list">
              <div v-for="(password, username) in defaultUsers" :key="username" class="credential-item">
                <strong>{{ username }}</strong> / {{ password }}
              </div>
            </div>
            <button @click="showCredentials = false" class="hide-btn">إخفاء</button>
          </div>
          <button v-else @click="showCredentials = true" class="show-credentials-btn">
            عرض بيانات تسجيل الدخول
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const formData = ref({
  username: '',
  password: ''
})

const loading = ref(false)
const errorMessage = ref('')
const showCredentials = ref(false)

// بيانات تسجيل الدخول الافتراضية
const defaultUsers = {
  'admin': 'admin123',
  'مدير': 'admin123',
  'user': 'user123',
  'مستخدم': 'user123',
  'ahmed': 'ahmed123',
  'أحمد': 'ahmed123',
  'fatima': 'fatima123',
  'فاطمة': 'fatima123',
  'mohammed': 'mohammed123',
  'محمد': 'mohammed123'
}

const handleLogin = async () => {
  // التحقق من الحقول
  if (!formData.value.username || !formData.value.password) {
    errorMessage.value = 'يرجى إدخال اسم المستخدم وكلمة المرور'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    // محاكاة تأخير للواقعية
    await new Promise(resolve => setTimeout(resolve, 800))
    
    // التحقق من بيانات تسجيل الدخول
    const username = formData.value.username.trim().toLowerCase()
    const password = formData.value.password
    
    // البحث في بيانات المستخدمين (دعم العربية والإنجليزية)
    const userKey = Object.keys(defaultUsers).find(
      key => key.toLowerCase() === username
    )
    
    if (userKey && defaultUsers[userKey] === password) {
      // حفظ حالة تسجيل الدخول
      localStorage.setItem('isAuthenticated', 'true')
      localStorage.setItem('username', formData.value.username)
      localStorage.setItem('loginTime', new Date().toISOString())
      
      // إعادة التوجيه إلى الصفحة الرئيسية
      await router.push('/')
    } else {
      errorMessage.value = 'اسم المستخدم أو كلمة المرور غير صحيحة'
    }
  } catch (error) {
    errorMessage.value = 'حدث خطأ أثناء تسجيل الدخول. يرجى المحاولة مرة أخرى.'
    console.error('Login error:', error)
  } finally {
    loading.value = false
  }
}

const handleForgotPassword = () => {
  // منطق استعادة كلمة المرور
  alert('ميزة استعادة كلمة المرور قريباً')
}

const handleRegister = () => {
  // منطق التسجيل
  alert('ميزة التسجيل قريباً')
}
</script>

<style scoped>
.login-page {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: var(--bg-main, #0a0a0a);
  padding: 20px;
  gap: 30px;
}

.system-title-wrapper {
  text-align: center;
  margin-bottom: 20px;
  z-index: 10;
  animation: fadeInDown 0.8s ease-out;
}

.system-title {
  color: var(--text-white, #ffffff);
  font-size: 2.5rem;
  font-weight: 700;
  margin: 0;
  text-shadow: 0 4px 12px rgba(0, 0, 0, 0.5);
  font-family: var(--font-main, 'Cairo', sans-serif);
  letter-spacing: 2px;
  background: linear-gradient(135deg, #ffffff 0%, #00f2ff 50%, #ffffff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  background-size: 200% auto;
  animation: gradient-shift 3s ease infinite;
}

@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes gradient-shift {
  0%, 100% {
    background-position: 0% center;
  }
  50% {
    background-position: 100% center;
  }
}

@property --a {
  syntax: '<angle>';
  initial-value: 0deg;
  inherits: false;
}

.box {
  position: relative;
  width: 400px;
  height: 200px;
  border-radius: 20px;
  background: repeating-conic-gradient(
    from var(--a),
    var(--primary, #00f2ff) 0%,
    var(--primary, #00f2ff) 5%,
    transparent 5%,
    transparent 40%,
    var(--primary, #00f2ff) 50%
  );
  animation: rotating 4s linear infinite;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: 0.5s;
}

.box:hover {
  width: 450px;
  height: 500px;
}

@keyframes rotating {
  to {
    --a: 360deg;
  }
}

.box::before {
  content: '';
  position: absolute;
  inset: 0;
  border-radius: 20px;
  background: repeating-conic-gradient(
    from var(--a),
    var(--secondary, #ff00ff) 0%,
    var(--secondary, #ff00ff) 5%,
    transparent 5%,
    transparent 40%,
    var(--secondary, #ff00ff) 50%
  );
  animation: rotating 4s linear infinite;
  animation-delay: -1s;
}

.box::after {
  content: '';
  position: absolute;
  inset: 4px;
  background: var(--box-bg, #1a1a1a);
  border-radius: 15px;
}

.login {
  position: absolute;
  inset: 60px;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;
}

.box:hover .login {
  inset: 40px;
}

.loginbx {
  width: 70%;
  display: flex;
  flex-direction: column;
  gap: 20px;
  transform: translateY(120px);
  transition: 0.5s;
  opacity: 0;
}

.box:hover .loginbx {
  transform: translateY(0);
  opacity: 1;
}

.loginbx h2 {
  color: var(--text-white, #ffffff);
  text-transform: uppercase;
  letter-spacing: 0.2em;
  text-align: center;
  font-size: 1.5rem;
  margin-bottom: 10px;
  font-family: var(--font-main, 'Cairo', sans-serif);
}

.loginbx h2 i {
  color: var(--primary, #00f2ff);
  margin-left: 10px;
}

.loginbx input {
  padding: 12px 20px;
  border-radius: 30px;
  border: none;
  outline: none;
  background: rgba(0, 0, 0, 0.3);
  color: var(--text-white, #ffffff);
  font-size: 1rem;
  font-family: var(--font-main, 'Cairo', sans-serif);
  transition: all 0.3s ease;
}

.loginbx input::placeholder {
  color: rgba(255, 255, 255, 0.5);
}

.loginbx input:focus {
  background: rgba(0, 0, 0, 0.5);
  box-shadow: 0 0 10px rgba(0, 242, 255, 0.3);
}

.loginbx input[type="submit"] {
  background: var(--secondary, #ff00ff);
  color: #111;
  cursor: pointer;
  font-weight: 600;
  margin-top: 10px;
  transition: all 0.3s ease;
}

.loginbx input[type="submit"]:hover:not(:disabled) {
  box-shadow: 0 0 20px var(--secondary, #ff00ff);
  transform: translateY(-2px);
}

.loginbx input[type="submit"]:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.loginbx input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error-message {
  background: rgba(244, 67, 54, 0.2);
  color: #ff5252;
  padding: 10px 15px;
  border-radius: 8px;
  font-size: 0.9rem;
  text-align: center;
  margin-bottom: 10px;
  border: 1px solid rgba(244, 67, 54, 0.3);
  font-family: var(--font-main, 'Cairo', sans-serif);
}

.group {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 10px;
}

.group a {
  color: var(--text-white, #ffffff);
  text-decoration: none;
  font-size: 0.9rem;
  transition: all 0.3s ease;
  font-family: var(--font-main, 'Cairo', sans-serif);
}

.group a:hover {
  text-decoration: underline;
}

.group a:last-child {
  color: var(--primary, #00f2ff);
}

.group a:last-child:hover {
  color: var(--secondary, #ff00ff);
}

/* معلومات بيانات تسجيل الدخول */
.credentials-info {
  margin-top: 20px;
  padding: 15px;
  background: rgba(0, 0, 0, 0.4);
  border-radius: 10px;
  border: 1px solid rgba(0, 242, 255, 0.3);
}

.credentials-info h3 {
  color: var(--primary, #00f2ff);
  font-size: 0.9rem;
  margin-bottom: 10px;
  text-align: center;
  font-family: var(--font-main, 'Cairo', sans-serif);
}

.credentials-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 10px;
  max-height: 150px;
  overflow-y: auto;
}

.credential-item {
  color: var(--text-white, #ffffff);
  font-size: 0.85rem;
  padding: 5px 10px;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 5px;
  font-family: var(--font-main, 'Cairo', sans-serif);
  text-align: center;
}

.credential-item strong {
  color: var(--primary, #00f2ff);
}

.hide-btn,
.show-credentials-btn {
  width: 100%;
  padding: 8px 15px;
  border-radius: 20px;
  border: 1px solid var(--primary, #00f2ff);
  background: rgba(0, 242, 255, 0.1);
  color: var(--primary, #00f2ff);
  cursor: pointer;
  font-size: 0.85rem;
  transition: all 0.3s ease;
  font-family: var(--font-main, 'Cairo', sans-serif);
  margin-top: 10px;
}

.hide-btn:hover,
.show-credentials-btn:hover {
  background: rgba(0, 242, 255, 0.2);
  box-shadow: 0 0 10px rgba(0, 242, 255, 0.3);
}

/* تحسينات للشاشات الصغيرة */
@media (max-width: 600px) {
  .box {
    width: 90%;
    max-width: 400px;
  }
  
  .box:hover {
    width: 95%;
    max-width: 450px;
  }
  
  .loginbx {
    width: 85%;
  }
  
  .loginbx h2 {
    font-size: 1.2rem;
  }
}
</style>

