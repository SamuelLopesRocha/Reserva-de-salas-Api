
from config import db

def getAluno():
    alunos = Aluno.query.all()
    return jsonify([aluno.to_dict() for aluno in alunos])