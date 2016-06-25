#ifndef GOMODELINTERFACE_H
#define GOMODELINTERFACE_H

#include <QAbstractListModel>

class Q_CORE_EXPORT GoModelInterface : public QAbstractListModel
{
    Q_OBJECT

public:
    explicit GoModelInterface(QObject *parent = 0);

    virtual int rowCount(const QModelIndex &parent = QModelIndex()) const;
    virtual QVariant data(const QModelIndex &index, int role) const;
    virtual QHash<int,QByteArray> roleNames() const;
};

#endif
