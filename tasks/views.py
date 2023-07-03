from django.shortcuts import render, redirect, get_object_or_404
from .models import Task

def task_list(request):
    tasks = Task.objects.all()
    return render(request, 'tasks/task_list.html', {'tasks': tasks})

def task_create(request):
    if request.method == 'POST':
        title = request.POST.get('title')
        Task.objects.create(title=title)
        return redirect('task_list')
    return render(request, 'tasks/task_create.html')

def task_delete(request):
    if request.method == 'POST':
        selected_tasks = request.POST.getlist('selected_tasks')
        Task.objects.filter(pk__in=selected_tasks).delete()
    return redirect('task_list')
