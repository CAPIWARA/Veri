pragma solidity ^0.4.6;

contract DocumentRegister {
    address public _address;
    mapping (string => Document) _documents;
    mapping(uint256 => Trajectory) _trajectory;

    struct Document {
        uint balance;
        uint256 insertAt;
        uint256[] trajectory;
    }

    struct Trajectory {
        string uuid;
        uint km;
        string trajectory;
        uint256 insertAt;
    }

    uint public numberOfDocuments = 0;

    function DocumentRegister() public {
        _address = msg.sender;
    }

    function DestroyContract() public {
        selfdestruct(_address);
    }

    function registryTrajectory(uint256 _points, string _trajeto, string _uuid, uint _km) payable public {
        numberOfDocuments = numberOfDocuments + 1; 
        Document memory contra = _documents[_uuid];

        Trajectory memory traj = Trajectory ({
            uuid: _uuid,
            insertAt: now,
            trajectory: _trajeto,
            km: _km
        });
        _trajectory[numberOfDocuments] = traj;

        uint256[] memory sub = new uint256[](contra.trajectory.length+1);
        sub[contra.trajectory.length] = traj.insertAt;

        Document memory doc = Document({
            balance: contra.balance + _points,
            trajectory: sub,
            insertAt: now
        });

        _documents[_uuid] = doc;
    }

    function consultByUuid(string _uuid) public view returns (uint256 balance, uint256[] trajectory) {
        Document memory doc = _documents[_uuid];
        return (doc.balance, doc.trajectory);
    }
}



