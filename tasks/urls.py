from django.urls import path, reverse
from . import views

urlpatterns = [
        path('', views.task_list, name='task_list'),
        path('create/', views.task_create, name='task_create'),
        path('tasks/delete/', views.task_delete, name='task_delete'),
        ]
