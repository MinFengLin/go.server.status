<!DOCTYPE html>
<html lang="en">
<head>
    <title>民有路 - 各項服務狀態</title>
    <meta content="text/html" charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="icon" href="https://cdn-icons-png.flaticon.com/512/2786/2786366.png">

    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">
    <style>
pre {
    overflow-x: auto;
    max-width: 60vw;
}

pre code {
    word-wrap: normal;
    white-space: pre;
}
    </style>
</head>
<html><div class="container">
<!-- Wi-Fi QR-Code -->
<div class="card my-2">
      <h4 class="card-header text-center">
        Wi-Fi QR-Code
      </h4>
    <div class="card mb-3 mx-auto" style="max-width: 540px;">
        <div class="row align-items-center g-0">
            <div class="col-md-4">
              <img src="Lin-qrcode.png" alt="Lin-home_Wi-Fi" class="img-thumbnail">
            </div>
            <div class="col-md-8">
                <div class="card-body">
                    <h5 class="card-title">網路名稱: </h5>
                    <p class="card-text">Lin</p>
                    <hr>
                    <h5 class="card-title">密碼: </h5>
                    <p class="card-text">26277050</p>
                </div>
            </div>
            <hr>
            <div class="col-md-12">
                <span class="d-inline-block"><small class="text-muted"><strong>將手機相機對準 QR Code 即可自動連接 WiFi</strong></small>
                </span>
                <i class="fas fa-angle-down">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-camera-fill" viewBox="0 0 16 16">
                        <path d="M10.5 8.5a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0z"/>
                        <path d="M2 4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-1.172a2 2 0 0 1-1.414-.586l-.828-.828A2 2 0 0 0 9.172 2H6.828a2 2 0 0 0-1.414.586l-.828.828A2 2 0 0 1 3.172 4H2zm.5 2a.5.5 0 1 1 0-1 .5.5 0 0 1 0 1zm9 2.5a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0z"/>
                    </svg>
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-phone-fill" viewBox="0 0 16 16">
                          <path d="M3 2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V2zm6 11a1 1 0 1 0-2 0 1 1 0 0 0 2 0z"/>
                    </svg>
                </i>
            </div>
        </div>
    </div>
</div>
<!-- End - Wi-Fi QR-Code -->
<?php
/*
 *
 * @author      Trường An Phạm Nguyễn
 * @copyright   2019, The authors
 * @license     GNU AFFERO GENERAL PUBLIC LICENSE
 *        http://www.gnu.org/licenses/agpl-3.0.html
 *
 * Jul 27, 2013
Original author:
*       Disclaimer Notice(s)                                                          
*       ex: This code is freely given to you and given "AS IS", SO if it damages      
*       your computer, formats your HDs, or burns your house I am not the one to
*       blame.                                                                     
*       Moreover, don't forget to include my copyright notices and name.               
*   +------------------------------------------------------------------------------+
*       Author(s): Crooty.co.uk (Adam C)                                    
*   +------------------------------------------------------------------------------+
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

$IoTdata = "";
$IoTdata .= '
<div class="card my-2">
  <h4 class="card-header text-center">
    IoT Service Status & Control
  </h4>
  <div class="card-body pb-0">
';

$IoTservices = array();

$IoTservices[] = array("service" => "加壓馬達", "ip" => "10.125.6.175");
$IoTdata .= "<small><table  class='table table-striped table-sm '><thead><tr><th>Service</th><th>Status</th><th>Control</th></tr></thead>";

foreach ($IoTservices as $service) {
    $pump_url = 'http://' . $service['ip'] . '/cm?cmnd=Power';
    $pump_json = file_get_contents($pump_url);
    $pump_json_decode = json_decode($pump_json);
    
    if($pump_json_decode->POWER == "OFF") {
        // $IoTdata .= "<tr><td>" . $service['service'] . "</td><td><a class='btn btn-primary' role='button' href=\"http://" . $service['ip'] . "/cm?cmnd=Power%20On\">Open</a></td>";
        // $IoTdata .= "<td class='table-danger'>Offline</td></tr>";
        $IoTdata .= "<tr><td><a href=\"http://" . $service['ip'] . "\"><strong>" . $service['service'] . "</strong></a></td><td class='table-danger'><strong>Shutdown</strong></td>";
        $IoTdata .= "<td><a class='btn btn-primary' role='button' href=\"http://" . $service['ip'] . "/cm?cmnd=Power%20On\">Open</a></td></tr>";
    } else {
        // $IoTdata .= "<tr><td>" . $service['service'] . "</td><td><a class='btn btn-danger' role='button' href=\"http://" . $service['ip'] . "/cm?cmnd=Power%20Off\">Close</a></td>";
        // $IoTdata .= "<td class='table-success'>Online</td></tr>";
        $IoTdata .= "<tr><td><a href=\"http://" . $service['ip'] . "\"><strong>" . $service['service'] . "</strong></a></td><td class='table-success'><strong>Working</strong></td>";
        $IoTdata .= "<td><a class='btn btn-danger' role='button' href=\"http://" . $service['ip'] . "/cm?cmnd=Power%20Off\">Close</a></td></tr>";
    }

}
//close table
$IoTdata .= "</table></small>";
$IoTdata .= '
  </div>
</div>
';
echo $IoTdata;

$data = "";
$data .= '
<div class="card my-2">
  <h4 class="card-header text-center">
    Service status
  </h4>
  <div class="card-body pb-0">
';


//configure script
$timeout = "1";

//set service checks
/* 
The script will open a socket to the following service to test for connection.
Does not test the fucntionality, just the ability to connect
Each service can have a name, port and the Unix domain it run on (default to localhost)
*/
$services = array();

$services[] = array("port" => "80",   "service" => "Mikrotik-RB750Gr3",                            "ip" => "10.125.6.254", "other_cfg" => "") ;
$services[] = array("port" => "80",   "service" => "HPE PoE Managed Switch",                       "ip" => "10.125.6.253", "other_cfg" => "") ;
$services[] = array("port" => "80",   "service" => "4F AP: RT-n18u Asuswrt-Merlin",                "ip" => "10.125.6.252", "other_cfg" => "") ;
$services[] = array("port" => "80",   "service" => "2F AP: Xiaomi Mi AX3600",                      "ip" => "10.125.6.251", "other_cfg" => "") ;
$services[] = array("port" => "80",   "service" => "HIKVISION 監視器",                             "ip" => "10.125.6.246", "other_cfg" => "") ;
$services[] = array("port" => "7050", "service" => "監視器影片存檔區 (KVM - SurveillanceStation)", "ip" => "",             "other_cfg" => "webman/3rdparty/SurveillanceStation") ;
$services[] = array("port" => "7050", "service" => "Synology NAS DSM",                             "ip" => "",             "other_cfg" => "") ;
$services[] = array("port" => "3000", "service" => "Gitea",                                        "ip" => "",             "other_cfg" => "") ;
$services[] = array("port" => "80",   "service" => "IoT: 加壓馬達",                                "ip" => "10.125.6.175", "other_cfg" => "") ;
$services[] = array("port" => "53",   "service" => "(Google DNS)Internet Connection",              "ip" => "8.8.8.8",      "other_cfg" => "") ;

http://10.125.6.250:7050/webman/3rdparty/SurveillanceStation/

//begin table for status
$data .= "<small><table  class='table table-striped table-sm '><thead><tr><th>Service</th><th>Port</th><th>Status</th></tr></thead>";
foreach ($services  as $service) {
    if($service['ip']=="" || $service['ip']=="localhost"){
       $service['ip'] = "10.125.6.250";
    }
    if($service['port']=="80" || $service['port']=="7050" || $service['port']=="3000"){
        if($service['other_cfg']==""){
            $data .= "<tr><td><a href=\"http://" . $service['ip'] . ":" . $service['port'] . "\">" . $service['service'] . "</a></td><td>". $service['port'];
        } else {
            $data .= "<tr><td><a href=\"http://" . $service['ip'] . ":" . $service['port'] . "/" . $service['other_cfg'] . "\">" . $service['service'] . "</a></td><td>". $service['port'];
        }
    } else {
        $data .= "<tr><td>" . $service['service'] . "</td><td>". $service['port'];
    }

    $fp = @fsockopen($service['ip'], $service['port'], $errno, $errstr, $timeout);
    if (!$fp) {
        $data .= "</td><td class='table-danger'>Offline </td></tr>";
      //fclose($fp);
    } else {
        $data .= "</td><td class='table-success'>Online</td></tr>";
        fclose($fp);
    }

}  
//close table
$data .= "</table></small>";
$data .= '
  </div>
</div>
';
echo $data;

// parser json fild
/*
 * json type example:
 * {
 *      "battery":{
 *          "charge":{
 *              "value":"100"
 *          },
 *          "runtime":{
 *              "value":"4530"
 *          }
 *      }
 *  }
 */
$Json = file_get_contents("ups_data.json");
// Converts to an array 
$ups_info_array = json_decode($Json);
// var_dump($ups_info_array);                     // prints ups_info_array
// echo $ups_info_array->battery->charge->value;  // prints charge value
// echo $ups_info_array->battery->runtime->value; // prints ups_info_array
// end - parser json fild

// Get charge / Get disk free
function getSymbolByQuantity($bytes) {
    $symbol = array('B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB');
    $exp = floor(log($bytes)/log(1024));
    
    return sprintf('%.2f<small>'.$symbol[$exp].'</small>', ($bytes/pow(1024, floor($exp))));
}
function percent_to_color($p){
    if($p < 30) return 'success';
    if($p < 45) return 'info';
    if($p < 60) return 'primary';
    if($p < 75) return 'warning';
    return 'danger';
}
function percent_to_color_charge($p){
    if($p < 30) return 'danger';
    if($p < 45) return 'warning';
    if($p < 60) return 'primary';
    if($p < 75) return 'info';
    return 'success';
}
function format_storage_info($disk_space, $disk_free, $disk_name, $battery_type){
    $str = "";
    $disk_free_precent = 100 - round($disk_free*1.0 / $disk_space*100, 2);
        $str .= '<div class="col p-0 d-inline-flex">';
        if ($disk_name == "charge") {
            $str .= "<span class='mr-2'>". badge($battery_type, 'secondary') ."</span>";
        } else {
            $str .= "<span class='mr-2'>" . badge($disk_name, 'secondary') .' '. getSymbolByQuantity($disk_free) . '/'. getSymbolByQuantity($disk_space) ."</span>";
        }
        $str .= '
<div class="progress flex-grow-1 align-self-center">
  <div class="progress-bar progress-bar-striped progress-bar-animated ';
        if ($disk_name == "charge") {
            $str .= 'bg-' . percent_to_color_charge($disk_free) .'
  " role="progressbar" style="width: '.$disk_free.'%;" aria-valuenow="'.$disk_free.'" aria-valuemin="0" aria-valuemax="100">'.$disk_free.'%</div>
</div>
</div>              ';
        } else {
            $str .= 'bg-' . percent_to_color($disk_free_precent) .'
  " role="progressbar" style="width: '.$disk_free_precent.'%;" aria-valuenow="'.$disk_free_precent.'" aria-valuemin="0" aria-valuemax="100">'.$disk_free_precent.'%</div>
</div>
</div>		        ';
        }
    return $str;

}
// end - Get charge / Get disk free

$ups_configs = array();

$ups_configs[] = array("config" => "status",          "config_name" => "狀態",       "value" => $ups_info_array->ups->status->value,         "unit" => "");
$ups_configs[] = array("config" => "charge",          "config_name" => "電池電量",   "value" => $ups_info_array->battery->charge->value,     "unit" => "%");
$ups_configs[] = array("config" => "runtime",         "config_name" => "預估時間",   "value" => $ups_info_array->battery->runtime->value,    "unit" => "s");
$ups_configs[] = array("config" => "battery_voltage", "config_name" => "電池電壓",   "value" => $ups_info_array->battery->voltage->value,    "unit" => "V");
$ups_configs[] = array("config" => "input_voltage",   "config_name" => "輸入電壓",   "value" => $ups_info_array->input->voltage->value,      "unit" => "V");
$ups_configs[] = array("config" => "beeper",          "config_name" => "蜂鳴器",     "value" => $ups_info_array->ups->beeper->status->value, "unit" => "");
$ups_configs[] = array("config" => "load",            "config_name" => "負載",       "value" => $ups_info_array->ups->load->value,           "unit" => "%");
$ups_configs[] = array("config" => "mfr",             "config_name" => "廠牌",       "value" => $ups_info_array->ups->mfr->value,            "unit" => "");
$ups_configs[] = array("config" => "test",            "config_name" => "電池測試",   "value" => $ups_info_array->ups->test->result->value,   "unit" => "");

$battery_charging_emoji = "<svg xmlns='http://www.w3.org/2000/svg' width='16' height='16' fill='currentColor' class='bi bi-battery-charging' viewBox='0 0 16 16'><path d='M9.585 2.568a.5.5 0 0 1 .226.58L8.677 6.832h1.99a.5.5 0 0 1 .364.843l-5.334 5.667a.5.5 0 0 1-.842-.49L5.99 9.167H4a.5.5 0 0 1-.364-.843l5.333-5.667a.5.5 0 0 1 .616-.09z'/><path d='M2 4h4.332l-.94 1H2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2.38l-.308 1H2a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2z'/><path d='M2 6h2.45L2.908 7.639A1.5 1.5 0 0 0 3.313 10H2V6zm8.595-2-.308 1H12a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H9.276l-.942 1H12a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-1.405z'/><path d='M12 10h-1.783l1.542-1.639c.097-.103.178-.218.241-.34V10zm0-3.354V6h-.646a1.5 1.5 0 0 1 .646.646zM16 8a1.5 1.5 0 0 1-1.5 1.5v-3A1.5 1.5 0 0 1 16 8z'/></svg>";
$battery_empty_emoji = "<svg xmlns='http://www.w3.org/2000/svg' width='16' height='16' fill='currentColor' class='bi bi-battery' viewBox='0 0 16 16'><path d='M0 6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V6zm2-1a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H2zm14 3a1.5 1.5 0 0 1-1.5 1.5v-3A1.5 1.5 0 0 1 16 8z'/></svg>";
$battery_full_emoji = "<svg xmlns='http://www.w3.org/2000/svg' width='16' height='16' fill='currentColor' class='bi bi-battery-full' viewBox='0 0 16 16'><path d='M2 6h10v4H2V6z'/><path d='M2 4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H2zm10 1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h10zm4 3a1.5 1.5 0 0 1-1.5 1.5v-3A1.5 1.5 0 0 1 16 8z'/></svg>";
$battery_half_emoji = "<svg xmlns='http://www.w3.org/2000/svg' width='16' height='16' fill='currentColor' class='bi bi-battery-half' viewBox='0 0 16 16'><path d='M2 6h5v4H2V6z'/><path d='M2 4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H2zm10 1a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h10zm4 3a1.5 1.5 0 0 1-1.5 1.5v-3A1.5 1.5 0 0 1 16 8z'/></svg>";
$battery_charging = false;

$data_ups = "";
$data_ups .= '
<div class="card mb-2">
  <h4 class="card-header text-center">
    UPS information
  </h4>
<div class="card-body">
';

$data_ups .= "<table  class='table table-sm mb-0'>";
foreach ($ups_configs as $ups_config) {
    if ($ups_config['config']=="status") {
        if ($ups_config['value']=="OL" || $ups_config['value']=="OL CHRG") $battery_charging = true;
        if ($ups_config['value']=="OL") {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td><span class='badge badge-success'>" .$ups_config['value']. "</span> 穩定在線 - 已充飽</td></tr>";
        } elseif ($ups_config['value']=="OL CHRG") {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td><span class='badge badge-info'>" .$ups_config['value']. "</span> 穩定在線 - 充電中</td></tr>";
        } elseif ($ups_config['value']=="OB DISCHRG") {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td><span class='badge badge-primary'>" .$ups_config['value']. "</span> 電池供應模式</td></tr>";
        } elseif ($ups_config['value']=="LB") {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td><span class='badge badge-warning'>" .$ups_config['value']. "</span> 電池供應模式 - 低電量</td></tr>";
        } else {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td><span class='badge badge-danger'>" .$ups_config['value']. "</span> 關閉中</td></tr>";
        }
    } elseif ($ups_config['config']=="charge") {
        if ($battery_charging == true) {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . " " . $battery_charging_emoji . "</td><td>". format_storage_info('', $ups_config['value'], $ups_config['config'], $ups_info_array->battery->type->value) ."</td></tr>";
        } else {
            if ($ups_config['value'] > 75) {
                $data_ups .= "<tr><td>" . $ups_config['config_name'] . " " . $battery_full_emoji . "</td><td>". format_storage_info('', $ups_config['value'], $ups_config['config'], $ups_info_array->battery->type->value) ."</td></tr>";
            } elseif ($ups_config['value'] > 35) {
                $data_ups .= "<tr><td>" . $ups_config['config_name'] . " " . $battery_half_emoji . "</td><td>". format_storage_info('', $ups_config['value'], $ups_config['config'], $ups_info_array->battery->type->value) ."</td></tr>";
            } else {
                $data_ups .= "<tr><td>" . $ups_config['config_name'] . " " . $battery_empty_emoji . "</td><td>". format_storage_info('', $ups_config['value'], $ups_config['config'], $ups_info_array->battery->type->value) ."</td></tr>";
            }
        }
    } elseif ($ups_config['config']=="mfr") {
        if ($ups_config['value']=="CPS") {
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td>CyberPower</td></tr>";
        } else {
            // TBC
            $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td>". $ups_config['value'] . " " .$ups_config['unit']. "</td></tr>";
        }
    } else {
        $data_ups .= "<tr><td>" . $ups_config['config_name'] . "</td><td>". $ups_config['value'] . " " .$ups_config['unit']. "</td></tr>";
    }
}
$data_ups .= "</table>";
$data_ups .= '  </div></div>';
// $data_ups .= '  </div>';
echo $data_ups;

/* =====================================================================
//
// ////////////////// SERVER INFORMATION  /////////////////////////////////
//
//
* =======================================================================/*/

$data1 = "";
$data1 .= '
<div class="card mb-2">
  <h4 class="card-header text-center">
    Synology NAS information
  </h4>
  <div class="card-body">
';

$data1 .= "<table  class='table table-sm mb-0'>";
// $data1 .= "<div class='table-responsive'><table  class='table table-sm mb-0'>";

//GET SERVER LOADS
$loadresult = @exec('uptime');  
preg_match("/averages?: ([0-9\.]+),[\s]+([0-9\.]+),[\s]+([0-9\.]+)/",$loadresult,$avgs);


//GET SERVER UPTIME
$uptime = explode(' up ', $loadresult);
$uptime = explode(',', $uptime[1]);
$uptime = $uptime[0].', '.$uptime[1];

function get_disk_free_status($disks){
    $str="";
    $max = 5;
    foreach($disks as $disk){
        if(strlen($disk["name"]) > $max) 
            $max = strlen($disk["name"]);
    }
    
    foreach($disks as $disk){
        $disk_space = disk_total_space($disk["path"]);
        $disk_free = disk_free_space($disk["path"]);

        $str .= format_storage_info($disk_space, $disk_free, $disk['name'], '');

    }
    return $str;
}
function badge($str, $type){
    return "<span class='badge badge-" . $type . " ' >$str</span>";
}

//Get ram usage
$total_mem = preg_split('/ +/', @exec('grep MemTotal /proc/meminfo'));
$total_mem = $total_mem[1];
$free_mem = preg_split('/ +/', @exec('grep MemFree /proc/meminfo'));
$cache_mem = preg_split('/ +/', @exec('grep ^Cached /proc/meminfo'));

$free_mem = $free_mem[1] + $cache_mem[1];

// //Get disk info
// $hdd_name= @exec(`smartctl -a /dev/hda -d ata | grep "Device Model" | awk '{print $3}'`);
// $hdd_temp = @exec(`smartctl -a /dev/hda -d ata | grep "Airflow_Temperature_Cel" | awk '{print $10}'`);

// echo $hdd_name;
// echo $hdd_temp;
// echo "123";

//Get top mem usage
$tom_mem_arr = array();
$top_cpu_use = array();

//-- The number of processes to display in Top RAM user
$i = 5;


/* ps command:
-e to display process from all user
-k to specify sorting order: - is desc order follow by column name
-o to specify output format, it's a list of column name. = suppress the display of column name
head to get only the first few lines 
*/
exec("ps -e k-rss -o rss,args | head -n $i | tail -n +2", $tom_mem_arr, $status);
exec("ps -e k-pcpu -o pcpu,args | head -n $i | tail -n +2", $top_cpu_use, $status);


$top_mem = implode('<br/>', $tom_mem_arr );
$top_mem = "<pre class='mb-0 '><code>" . $top_mem . "</code></pre>";

$top_cpu = implode('<br/>', $top_cpu_use );
$top_cpu = "<pre class='mb-0 '><code>" . $top_cpu. "</code></pre>";

$data1 .= "<tr><td>Average load</td><td><h5>". badge($avgs[1],'secondary'). ' ' .badge($avgs[2], 'secondary') . ' ' . badge( $avgs[3], 'secondary') . " </h5></td>\n";
$data1 .= "<tr><td>Uptime</td><td>$uptime                     </td></tr>";


$disks = array();

/*
* The disks array list all mountpoint you wan to check freespace
* Display name and path to the moutpoint have to be provide, you can 
*/
$disks[] = array("name" => "local" , "path" => getcwd()) ;
// $disks[] = array("name" => "Your disk name" , "path" => '/mount/point/to/that/disk') ;


$data1 .= "<tr><td>Disk free        </td><td>" . get_disk_free_status($disks) . "</td></tr>";

$data1 .= "<tr><td>RAM free        </td><td>". format_storage_info($total_mem *1024, $free_mem *1024, '', '') ."</td></tr>";
$data1 .= "<tr><td>Top RAM user    </td><td><small><pre class='mb-0 '><code>RSS COMMAND</code></pre>$top_mem</small></td></tr>";
$data1 .= "<tr><td>Top CPU user    </td><td><small><pre class='mb-0 '><code>CPU COMMAND</code></pre>$top_cpu</small></td></tr>";

$data1 .= "</table>";
$data1 .= '  </div></div>';
// $data1 .= '  </div>';
echo $data1;

/* =============================================================================
*
* DISPLAY BANDWIDTH STATISTIC, REQUIRE VNSTAT INSTALLED AND PROPERLY CONFIGURED.
*
* ===============================================================================s
*/

$traffic_default = "d";

$data2 = "";
$data2 .=  '
<div class="card mb-2">
  <h4 class="card-header text-center">
    vnstat Network traffic
  </h4>
  <div class="card-body text-center">
';


$data2 .="<span class=' d-block'><pre class='d-inline-block text-left'><small>";
$traffic_arr = array();
if (!isset($_GET['showtraffic']) || $_GET['showtraffic'] ==  false) {
    exec('/opt/bin/vnstat -' . $traffic_default, $traffic_arr, $status);
} else {
    exec('/opt/bin/vnstat -' . escapeshellarg( $_GET['showtraffic'] ), $traffic_arr, $status);
}

///for testing
// $traffic = "
// enp0s20  /  monthly
// month        rx      |     tx      |    total    |   avg. rate
// ------------------------+-------------+-------------+---------------
// Sep '18     36.60 GiB |    7.04 GiB |   43.64 GiB |  144.62 kbit/s
// Oct '18    400.69 GiB |    1.19 TiB |    1.58 TiB |    5.19 Mbit/s
// Nov '18    393.52 GiB |    2.19 TiB |    2.57 TiB |    8.72 Mbit/s
// Dec '18    507.28 GiB |    2.05 TiB |    2.55 TiB |    8.37 Mbit/s
// Jan '19    269.01 GiB |    1.39 TiB |    1.65 TiB |    7.51 Mbit/s
// ------------------------+-------------+-------------+---------------
// estimated    371.92 GiB |    1.92 TiB |    2.29 TiB |
// ";
/// for real
$traffic = implode("\n", $traffic_arr);

$data2 .="$traffic</small></pre></span>";

echo $data2;
?>
</div></div></html>