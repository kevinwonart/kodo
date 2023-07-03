from django.db import models

class Task(models.Model):
    id = models.AutoField(primary_key=True)

    title = models.CharField(max_length=202)
    completed = models.BooleanField(default=False)


    def __str__(self):
        return self.title
