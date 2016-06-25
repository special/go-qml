#include "gomodelinterface.h"
#include "capi.h"

GoModelInterface::GoModelInterface(QObject *parent)
    : QAbstractListModel(parent)
{
}

int GoModelInterface::rowCount(const QModelIndex &parent) const
{
    return hookModelRowCount((void*)this);
}

QVariant GoModelInterface::data(const QModelIndex &index, int role) const
{
    DataValue value;
    hookModelData((void*)this, index.row(), role, &value);

    QVariant var;
    unpackDataValue(&value, &var);
    return var;
}

QHash<int,QByteArray> GoModelInterface::roleNames() const
{
    DataValue value;
    hookModelRoleNames((void*)this, &value);

    QVariant var;
    unpackDataValue(&value, &var);

    // This is a horrible hack, because we can't return lists or maps from Go
    QHash<int,QByteArray> re;
    QString str = var.toString();
    QStringList roles = str.split(';');
    foreach (const QString &s, roles) {
        if (s.isEmpty())
            break;
        QStringList v = s.split('=');
        re[v[0].toInt()] = v[1].toLatin1();
    }

    return re;
}

