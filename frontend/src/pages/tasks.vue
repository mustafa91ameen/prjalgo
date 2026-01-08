<template>
  <div class="tasks-container">
    <!-- Page Header -->
    <PageHeader
      title="إدارة المهام"
      subtitle="إدارة وتتبع جميع المهام"
      badge="المهام"
      badgeType="info"
      class="tasks-header"
    />

    <!-- Tasks Content -->
    <div class="tasks-content">
      <span class="tasks-content-title">قائمة المهام</span>
      <button class="add-task-btn" @click="showAddModal = true">
        <i class="mdi mdi-plus"></i>
        إضافة مهمة جديدة
      </button>
    </div>

    <!-- Tasks Table -->
    <div v-if="tasks.length > 0" class="tasks-table-container">
      <table class="tasks-table">
        <thead>
          <tr>
            <th>#</th>
            <th>عنوان المهمة</th>
            <th>الحالة</th>
            <th>الأولوية</th>
            <th>المسؤول</th>
            <th>التاريخ</th>
            <th>الإجراءات</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(task, index) in tasks" :key="index">
            <td>{{ index + 1 }}</td>
            <td>{{ task.title }}</td>
            <td>
              <span :class="['status-badge', getStatusClass(task.status)]">
                {{ getStatusLabel(task.status) }}
              </span>
            </td>
            <td>
              <span :class="['priority-badge', getPriorityClass(task.priority)]">
                {{ getPriorityLabel(task.priority) }}
              </span>
            </td>
            <td>{{ getAssigneeName(task.assignee) }}</td>
            <td>{{ task.date }}</td>
            <td>
              <div class="action-buttons">
                <button class="action-btn edit" title="تعديل">
                  <i class="mdi mdi-pencil"></i>
                </button>
                <button class="action-btn delete" @click="deleteTask(index)" title="حذف">
                  <i class="mdi mdi-delete"></i>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Empty State -->
    <div v-else class="empty-state">
      <i class="mdi mdi-clipboard-text-outline"></i>
      <p>لا توجد مهام حالياً</p>
      <span>اضغط على "إضافة مهمة جديدة" لإضافة مهمة</span>
    </div>

    <!-- Add Task Modal -->
    <div v-if="showAddModal" class="modal-overlay" @click.self="showAddModal = false">
      <div class="modal-container">
        <div class="modal-header">
          <h3>إضافة مهمة جديدة</h3>
          <button class="modal-close" @click="showAddModal = false">
            <i class="mdi mdi-close"></i>
          </button>
        </div>
        
        <div class="modal-body">
          <!-- عنوان المهمة -->
          <div class="form-group">
            <label for="taskTitle">عنوان المهمة</label>
            <input 
              type="text" 
              id="taskTitle" 
              v-model="newTask.title" 
              placeholder="أدخل عنوان المهمة"
            />
          </div>

          <!-- حالة المهمة -->
          <div class="form-group">
            <label for="taskStatus">حالة المهمة</label>
            <select id="taskStatus" v-model="newTask.status">
              <option value="">اختر الحالة</option>
              <option value="pending">قيد الانتظار</option>
              <option value="in-progress">قيد التنفيذ</option>
              <option value="completed">مكتملة</option>
              <option value="cancelled">ملغاة</option>
            </select>
          </div>

          <!-- وصف المهمة -->
          <div class="form-group">
            <label for="taskDescription">وصف المهمة</label>
            <textarea 
              id="taskDescription" 
              v-model="newTask.description" 
              placeholder="أدخل وصف المهمة"
              rows="3"
            ></textarea>
          </div>

          <!-- أولوية المهمة -->
          <div class="form-group">
            <label for="taskPriority">أولوية المهمة</label>
            <select id="taskPriority" v-model="newTask.priority">
              <option value="">اختر الأولوية</option>
              <option value="low">منخفضة</option>
              <option value="medium">متوسطة</option>
              <option value="high">عالية</option>
              <option value="urgent">عاجلة</option>
            </select>
          </div>

          <!-- المسؤول -->
          <div class="form-group">
            <label for="taskAssignee">المسؤول</label>
            <select id="taskAssignee" v-model="newTask.assignee">
              <option value="">اختر المسؤول</option>
              <option value="user1">أحمد محمد</option>
              <option value="user2">خالد علي</option>
              <option value="user3">سارة أحمد</option>
              <option value="user4">فاطمة حسن</option>
            </select>
          </div>

          <!-- تاريخ المهمة -->
          <div class="form-group">
            <label for="taskDate">تاريخ المهمة</label>
            <input 
              type="date" 
              id="taskDate" 
              v-model="newTask.date"
            />
          </div>
        </div>

        <div class="modal-footer">
          <button class="btn-cancel" @click="showAddModal = false">إلغاء</button>
          <button class="btn-save" @click="saveTask">حفظ المهمة</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import PageHeader from '../components/PageHeader.vue'

const showAddModal = ref(false)
const tasks = ref([])

const newTask = reactive({
  title: '',
  status: '',
  description: '',
  priority: '',
  assignee: '',
  date: ''
})

const saveTask = () => {
  if (!newTask.title) {
    alert('يرجى إدخال عنوان المهمة')
    return
  }
  
  // إضافة المهمة للقائمة
  tasks.value.push({ ...newTask })
  
  // إعادة تعيين الحقول
  Object.assign(newTask, {
    title: '',
    status: '',
    description: '',
    priority: '',
    assignee: '',
    date: ''
  })
  
  showAddModal.value = false
}

const deleteTask = (index) => {
  tasks.value.splice(index, 1)
}

const getStatusLabel = (status) => {
  const labels = {
    'pending': 'قيد الانتظار',
    'in-progress': 'قيد التنفيذ',
    'completed': 'مكتملة',
    'cancelled': 'ملغاة'
  }
  return labels[status] || status
}

const getStatusClass = (status) => {
  return `status-${status}`
}

const getPriorityLabel = (priority) => {
  const labels = {
    'low': 'منخفضة',
    'medium': 'متوسطة',
    'high': 'عالية',
    'urgent': 'عاجلة'
  }
  return labels[priority] || priority
}

const getPriorityClass = (priority) => {
  return `priority-${priority}`
}

const getAssigneeName = (assignee) => {
  const names = {
    'user1': 'أحمد محمد',
    'user2': 'خالد علي',
    'user3': 'سارة أحمد',
    'user4': 'فاطمة حسن'
  }
  return names[assignee] || assignee
}
</script>

<style scoped>
.tasks-container {
  padding: 32px;
  max-width: 1600px;
  margin: 0 auto;
}

/* Tasks Header Custom Color */
.tasks-header {
  background: linear-gradient(135deg, #018790 0%, #005461 100%) !important;
}

.tasks-header::before {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 50%, #06b6d4 100%) !important;
}

.tasks-content {
  margin-top: 24px;
  padding: 16px 24px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  border: 2px solid rgba(6, 182, 212, 0.3);
  color: rgba(255, 255, 255, 0.95);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tasks-content-title {
  font-size: 18px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.95);
}

.add-task-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  border: none;
  border-radius: 10px;
  color: white;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(6, 182, 212, 0.3);
}

.add-task-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(6, 182, 212, 0.4);
}

.add-task-btn i {
  font-size: 18px;
}

/* Tasks Table */
.tasks-table-container {
  margin-top: 24px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  border: 2px solid rgba(6, 182, 212, 0.3);
  overflow: hidden;
}

.tasks-table {
  width: 100%;
  border-collapse: collapse;
}

.tasks-table thead {
  background: rgba(6, 182, 212, 0.15);
}

.tasks-table th {
  padding: 16px;
  text-align: right;
  color: rgba(255, 255, 255, 0.95);
  font-weight: 600;
  font-size: 14px;
  border-bottom: 1px solid rgba(6, 182, 212, 0.2);
}

.tasks-table td {
  padding: 16px;
  text-align: right;
  color: rgba(255, 255, 255, 0.85);
  font-size: 14px;
  border-bottom: 1px solid rgba(6, 182, 212, 0.1);
}

.tasks-table tbody tr {
  transition: all 0.3s ease;
}

.tasks-table tbody tr:hover {
  background: rgba(6, 182, 212, 0.1);
}

.tasks-table tbody tr:last-child td {
  border-bottom: none;
}

/* Status Badges */
.status-badge {
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.status-pending {
  background: rgba(251, 191, 36, 0.2);
  color: #fbbf24;
}

.status-in-progress {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.status-completed {
  background: rgba(16, 185, 129, 0.2);
  color: #10b981;
}

.status-cancelled {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

/* Priority Badges */
.priority-badge {
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.priority-low {
  background: rgba(107, 114, 128, 0.2);
  color: #9ca3af;
}

.priority-medium {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.priority-high {
  background: rgba(249, 115, 22, 0.2);
  color: #f97316;
}

.priority-urgent {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  gap: 8px;
  justify-content: center;
}

.action-btn {
  width: 32px;
  height: 32px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.action-btn i {
  font-size: 16px;
}

.action-btn.edit {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}

.action-btn.edit:hover {
  background: rgba(59, 130, 246, 0.4);
}

.action-btn.delete {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

.action-btn.delete:hover {
  background: rgba(239, 68, 68, 0.4);
}

/* Empty State */
.empty-state {
  margin-top: 24px;
  padding: 60px 24px;
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 16px;
  border: 2px solid rgba(6, 182, 212, 0.3);
  text-align: center;
}

.empty-state i {
  font-size: 64px;
  color: rgba(6, 182, 212, 0.4);
  margin-bottom: 16px;
}

.empty-state p {
  color: rgba(255, 255, 255, 0.9);
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 8px 0;
}

.empty-state span {
  color: rgba(255, 255, 255, 0.5);
  font-size: 14px;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(5px);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.modal-container {
  background: linear-gradient(135deg, #0a3d42 0%, #052428 100%);
  border-radius: 20px;
  border: 2px solid rgba(6, 182, 212, 0.4);
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 25px 60px rgba(0, 0, 0, 0.5);
  animation: slideUp 0.3s ease;
}

@keyframes slideUp {
  from { 
    opacity: 0;
    transform: translateY(30px);
  }
  to { 
    opacity: 1;
    transform: translateY(0);
  }
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(6, 182, 212, 0.2);
}

.modal-header h3 {
  margin: 0;
  color: #fff;
  font-size: 20px;
  font-weight: 600;
}

.modal-close {
  background: rgba(255, 255, 255, 0.1);
  border: none;
  width: 36px;
  height: 36px;
  border-radius: 10px;
  color: rgba(255, 255, 255, 0.7);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.modal-close:hover {
  background: rgba(239, 68, 68, 0.3);
  color: #ef4444;
}

.modal-close i {
  font-size: 20px;
}

.modal-body {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  color: rgba(255, 255, 255, 0.9);
  font-size: 14px;
  font-weight: 500;
}

.form-group input,
.form-group select,
.form-group textarea {
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(6, 182, 212, 0.3);
  border-radius: 10px;
  padding: 12px 16px;
  color: #fff;
  font-size: 14px;
  transition: all 0.3s ease;
  outline: none;
}

.form-group input::placeholder,
.form-group textarea::placeholder {
  color: rgba(255, 255, 255, 0.4);
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  border-color: #06b6d4;
  background: rgba(6, 182, 212, 0.1);
  box-shadow: 0 0 0 3px rgba(6, 182, 212, 0.1);
}

.form-group select {
  cursor: pointer;
}

.form-group select option {
  background: #0a3d42;
  color: #fff;
}

.form-group textarea {
  resize: vertical;
  min-height: 80px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 20px 24px;
  border-top: 1px solid rgba(6, 182, 212, 0.2);
}

.btn-cancel,
.btn-save {
  padding: 12px 24px;
  border-radius: 10px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  border: none;
}

.btn-cancel {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
}

.btn-cancel:hover {
  background: rgba(255, 255, 255, 0.2);
}

.btn-save {
  background: linear-gradient(135deg, #06b6d4 0%, #10b981 100%);
  color: white;
  box-shadow: 0 4px 15px rgba(6, 182, 212, 0.3);
}

.btn-save:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(6, 182, 212, 0.4);
}

/* Responsive Styles */
@media (max-width: 1024px) {
  .tasks-container {
    padding: 24px;
  }

  .tasks-content {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .add-task-btn {
    width: 100%;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .tasks-container {
    padding: 16px;
  }

  .tasks-content {
    padding: 14px 18px;
  }

  .tasks-content-title {
    font-size: 16px;
  }

  .add-task-btn {
    padding: 8px 16px;
    font-size: 13px;
  }

  .add-task-btn i {
    font-size: 16px;
  }

  /* Table responsive */
  .tasks-table-container {
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }

  .tasks-table {
    min-width: 700px;
  }

  .tasks-table th,
  .tasks-table td {
    padding: 12px;
    font-size: 13px;
    white-space: nowrap;
  }

  .status-badge,
  .priority-badge {
    padding: 4px 10px;
    font-size: 11px;
  }

  .action-btn {
    width: 28px;
    height: 28px;
  }

  .action-btn i {
    font-size: 14px;
  }

  /* Modal responsive */
  .modal-container {
    width: 95%;
    max-height: 85vh;
  }

  .modal-header {
    padding: 16px 20px;
  }

  .modal-header h3 {
    font-size: 18px;
  }

  .modal-body {
    padding: 20px;
    gap: 16px;
  }

  .modal-footer {
    padding: 16px 20px;
  }

  /* Empty state */
  .empty-state {
    padding: 48px 20px;
  }

  .empty-state i {
    font-size: 56px;
  }

  .empty-state p {
    font-size: 16px;
  }

  .empty-state span {
    font-size: 13px;
  }
}

@media (max-width: 480px) {
  .tasks-container {
    padding: 12px;
  }

  .tasks-content {
    margin-top: 16px;
    padding: 12px 14px;
    border-radius: 12px;
  }

  .tasks-content-title {
    font-size: 14px;
  }

  .add-task-btn {
    padding: 8px 14px;
    font-size: 12px;
    border-radius: 8px;
  }

  .add-task-btn i {
    font-size: 14px;
  }

  /* Table */
  .tasks-table-container {
    margin-top: 16px;
    border-radius: 12px;
  }

  .tasks-table th,
  .tasks-table td {
    padding: 10px 8px;
    font-size: 12px;
  }

  .status-badge,
  .priority-badge {
    padding: 4px 8px;
    font-size: 10px;
    border-radius: 12px;
  }

  .action-btn {
    width: 26px;
    height: 26px;
    border-radius: 6px;
  }

  .action-btn i {
    font-size: 13px;
  }

  .action-buttons {
    gap: 6px;
  }

  /* Modal full width */
  .modal-container {
    width: calc(100% - 16px);
    max-width: none;
    max-height: 90vh;
    margin: 8px;
    border-radius: 16px;
  }

  .modal-header {
    padding: 14px 16px;
  }

  .modal-header h3 {
    font-size: 16px;
  }

  .modal-close {
    width: 32px;
    height: 32px;
    border-radius: 8px;
  }

  .modal-body {
    padding: 16px;
    gap: 14px;
  }

  .form-group label {
    font-size: 13px;
  }

  .form-group input,
  .form-group select,
  .form-group textarea {
    padding: 10px 14px;
    font-size: 13px;
    border-radius: 8px;
  }

  .modal-footer {
    padding: 14px 16px;
    flex-wrap: wrap;
    gap: 8px;
  }

  .btn-cancel,
  .btn-save {
    padding: 10px 18px;
    font-size: 13px;
    flex: 1;
    min-width: 100px;
    text-align: center;
    justify-content: center;
  }

  /* Empty state */
  .empty-state {
    margin-top: 16px;
    padding: 40px 16px;
    border-radius: 12px;
  }

  .empty-state i {
    font-size: 48px;
    margin-bottom: 12px;
  }

  .empty-state p {
    font-size: 14px;
  }

  .empty-state span {
    font-size: 12px;
  }
}

@media (max-width: 360px) {
  .tasks-container {
    padding: 8px;
  }

  .tasks-content {
    padding: 10px 12px;
  }

  .tasks-content-title {
    font-size: 13px;
  }

  .add-task-btn {
    padding: 6px 12px;
    font-size: 11px;
  }

  .tasks-table th,
  .tasks-table td {
    padding: 8px 6px;
    font-size: 11px;
  }

  .modal-header {
    padding: 12px 14px;
  }

  .modal-header h3 {
    font-size: 15px;
  }

  .modal-body {
    padding: 12px;
  }

  .modal-footer {
    padding: 12px 14px;
  }

  .btn-cancel,
  .btn-save {
    padding: 8px 14px;
    font-size: 12px;
  }
}
</style>
