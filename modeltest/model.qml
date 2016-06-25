import QtQuick 2.1

Item {
    width: 800
    height: 600

    ListView {
        anchors.fill: parent
        spacing: 10

        delegate: Rectangle {
            color: "black"
            width: parent.width
            height: 30

            Text {
                anchors.centerIn: parent
                color: "white"
                text: magic
            }
        }

        model: goModel
    }
}
