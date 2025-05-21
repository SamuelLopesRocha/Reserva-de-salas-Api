from config import db

class Sala(db.Model):
    __tablename__ = "salas"

    sala_id = db.Column(db.Integer, primary_key=True, autoincrement=True)
    recursos = db.Column(db.String(100), nullable=False)
    ativo = db.Column(db.Boolean, default=True)

    def __init__(self, recursos, ativo=True):
        self.recursos = recursos
        self.ativo = ativo

    def to_dict(self):
        return {
            'sala_id': self.sala_id,
            'recursos': self.recursos,
            'ativo': self.ativo,
        }

