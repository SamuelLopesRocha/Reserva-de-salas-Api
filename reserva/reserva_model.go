from datetime import datetime
from config import db

id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    nome = db.Column(db.String(50), nullable=False)
    idade = db.Column(db.Integer, nullable=False)  # Este campo pode ser preenchido pelo cálculo no momento da criação
    turma_id = db.Column(db.Integer, db.ForeignKey('turmas.id'), nullable=False)
    data_de_nascimento = db.Column(db.Date, nullable=False)
    nota_primeiro_semestre = db.Column(db.Float, nullable=False)
    nota_segundo_semestre = db.Column(db.Float, nullable=False)
    media_final = db.Column(db.Float, nullable=False)  # Calculado como a média das notas

turma = db.relationship('Turma', backref='alunos')