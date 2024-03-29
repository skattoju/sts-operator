/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type StsGnssSpec struct {
	//
	// Enable/disable GPS
	//
	// Valid range 0-1
	//
	//0 - Disable GPS
	//
	//1 - Enable GPS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGpsEn int `json:"gnssSigGpsEn"`

	//
	// Enable/disable GPS L1C/A
	// Valid range 0-1
	//
	//     0 - Disable GPS L1C/A
	//
	//     1 - Enable GPS L1C/A (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGpsL1CAEn int `json:"gnssSigGpsL1CAEn"`

	//
	// Enable/disable GPS L2C
	// Valid range 0-1
	//
	//     0 - Disable GPS L2C
	//
	//     1 - Enable GPS L2C (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGpsL2CEn int `json:"gnssSigGpsL2CEn"`

	//
	// Enable/disable SBAS
	// Valid range 0-1
	//
	//     0 - Disable SBAS
	//
	//     1 - Enable SBAS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigSBASEn int `json:"gnssSigSBASEn"`

	//
	// Enable/disable SBAS L1C/A
	// Valid range 0-1
	//
	//     0 - Disable SBAS L1C/A (default)
	//
	//     1 - Enable SBAS L1C/A
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigSBASL1CAEn int `json:"gnssSigSBASL1CAEn"`

	//
	// Enable/disable Galileo
	// Valid range 0-1
	//
	//     0 - Disable Galileo
	//
	//     1 - Enable Galileo (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGalEn int `json:"gnssSigGalEn"`

	//
	// Enable/disable Galileo E1
	// Valid range 0-1
	//
	//     0 - Disable Galileo E1
	//
	//     1 - Enable Galileo E1 (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGalE1En int `json:"gnssSigGalE1En"`

	//
	// Enable/disable Galileo E5b
	// Valid range 0-1
	//
	//     0 - Disable Galileo E5b
	//
	//     1 - Enable Galileo E5b (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGalE5BEn int `json:"gnssSigGalE5BEn"`

	//
	// Enable/disable BeiDou
	// Valid range 0-1
	//
	//     0 - Disable BeiDou
	//
	//     1 - Enable BeiDou (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigBDSEn int `json:"gnssSigBDSEn"`

	//
	// Enable/disable BeiDou B1I
	// Valid range 0-1
	//
	//     0 - Disable BeiDou B1I
	//
	//     1 - Enable BeiDou B1I (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigBDSB1En int `json:"gnssSigBDSB1En"`

	//
	// Enable/disable BeiDou B2I
	// Valid range 0-1
	//
	//     0 - Disable BeiDou B2I
	//
	//     1 - Enable BeiDou B2I (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigBDSB2En int `json:"gnssSigBDSB2En"`

	//
	// Enable/disable QZSS
	// Valid range 0-1
	//
	//     0 - Disable QZSS
	//
	//     1 - Enable QZSS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigQZSSEn int `json:"gnssSigQZSSEn"`

	//
	// Enable/disable QZSS L1C/A
	// Valid range 0-1
	//
	//     0 - Disable QZSS L1C/A
	//
	//     1 - Enable QZSS L1C/A (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigQZSSL1CAEn int `json:"gnssSigQZSSL1CAEn"`

	//
	// Enable/disable QZSS L1S
	// Valid range 0-1
	//
	//     0 - Disable QZSS L1S (default)
	//
	//     1 - Enable QZSS L1S
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigQZSSL1SEn int `json:"gnssSigQZSSL1SEn"`

	//
	// Enable/disable QZSS L2C
	// Valid range 0-1
	//
	//     0 - Disable QZSS L2C
	//
	//     1 - Enable QZSS L2C (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigQZSSL2CEn int `json:"gnssSigQZSSL2CEn"`

	//
	// Enable/disable GLONASS
	// Valid range 0-1
	//
	//     0 - Disable GLONASS
	//
	//     1 - Enable GLONASS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGLOEn int `json:"gnssSigGLOEn"`

	//
	// Enable/disable GLONASS L1
	// Valid range 0-1
	//
	//     0 - Disable GLONASS L1
	//
	//     1 - Enable GLONASS L1 (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGLOL1En int `json:"gnssSigGLOL1En"`

	//
	// Enable/disable GLONASS L2
	// Valid range 0-1
	//
	//     0 - Disable GLONASS L2
	//
	//     1 - Enable GLONASS L2 (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssSigGLOL2En int `json:"gnssSigGLOL2En"`

	//
	// ********************************* CFG-TP- Time Pulse Configuration ******************************
	//

	//
	// Antenna cable delay settings (nsec)
	// Valid range +-50,000,000
	//
	//     N - nanoseconds
	//
	// +kubebuilder:validation:Minimum=-50000000
	// +kubebuilder:validation:Maximum=50000000
	// +kubebuilder:validation:Optional
	GnssCableDelay int `json:"gnssCableDelay"`

	//
	// Time Pulse is interpreted as Period or Frequency
	// Valid range 0-1
	//
	//     0 - Period
	//
	//     1 - Frequency (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssPulseDef int `json:"gnssPulseDef"`

	//
	// Time Pulse Length is interpreted as Pulse Ratio (%) or Length (nsec)
	// Valid range 0-1
	//
	//     0 - Pulse Ratio (default)
	//
	//     1 - Length
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssPulseLenDef int `json:"gnssPulseLenDef"`

	//
	// Enable/disable the first Time Pulse
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssTP1En int `json:"gnssTP1En"`

	//
	// Set Time Pulse Frequency (Hz) for the first Time Pulse
	// Valid range
	//
	//     1 - (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssFreqTP1 int `json:"gnssFreqTP1"`

	//
	// Set Time pulse Frequency (Hz), when locked to GNSS time, for the first Time Pulse
	// Valid range
	//
	//     1 - (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssFreqLockTP1 int `json:"gnssFreqLockTP1"`

	//
	// Use locked parameters, when possible, for the first Time Pulse
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssUseLockTP1 int `json:"gnssUseLockTP1"`

	//
	// Set Time Pulse Duty Cycle (%) for the first Time Pulse
	// Valid range 0-100
	//
	//     0 - (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssDutyTP1 int `json:"gnssDutyTP1"`

	//
	// Set Time Pulse Duty Cycle (%), when locked to GNSS time, for the first Time Pulse
	// Valid range 0-100
	//
	//     10 - (default)
	//
	// +kubebuilder:default:=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Optional
	GnssDutyLockTP1 int `json:"gnssDutyLockTP1"`

	//
	// Enable/disable the second Time Pulse (10 MHz)
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssTP2En int `json:"gnssTP2En"`

	//
	// Set Time Pulse Frequency (Hz) for the second Time Pulse
	// Valid range
	//
	//     10000000 - (default)
	//
	// +kubebuilder:default:=10000000
	// +kubebuilder:validation:Optional
	GnssFreqTP2 int `json:"gnssFreqTP2"`

	//
	// Set Time pulse Frequency (Hz), when locked to GNSS time, for the second Time Pulse
	// Valid range
	//
	//     10000000 - (default)
	//
	// +kubebuilder:default:=10000000
	// +kubebuilder:validation:Optional
	GnssFreqLockTP2 int `json:"gnssFreqLockTP2"`

	//
	// Use locked parameters, when possible, for the second Time Pulse
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssUseLockTP2 int `json:"gnssUseLockTP2"`

	//
	// Set Time Pulse Duty Cycle (%) for the second Time Pulse
	// Valid range 0-100
	//
	//     0 - (default)
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Optional
	GnssDutyTP2 int `json:"gnssDutyTP2"`

	//
	// Set Time Pulse Duty Cycle (%), when locked to GNSS time, for the second Time Pulse
	// Valid range 0-100
	//
	//     50 - (default)
	//
	// +kubebuilder:default:=50
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Optional
	GnssDutyLockTP2 int `json:"gnssDutyLockTP2"`

	//
	// ********************** CFG-NAVSPG- Standard Precision Navigation Configuration ******************
	//

	//
	// Minimum elevation level (degrees)
	// Valid range
	//
	//     N - degrees
	//
	// +kubebuilder:validation:Optional
	GnssMinElev int `json:"gnssMinElev"`

	//
	// Minimum satellite signal level for navigation (dBHz)
	// Valid range
	//
	//     N - dBHz
	//
	// +kubebuilder:validation:Optional
	GnssMinSatSig int `json:"gnssMinSatSig"`

	//
	// ************** CFG-INFMSG- Information Message Configuration **************
	//

	//
	// Information Message Flags for NMEA Protocol on the USB (bitmask)
	// Valid range 0-31
	//
	//     Bit 0 - Enable (default) / Disable Error Information Messages
	//
	//     Bit 1 - Enable (default) / Disable Warning Information Messages
	//
	//     Bit 2 - Enable (default) / Disable Notice Information Messages
	//
	//     Bit 3 - Enable/Disable (default) Test Information Messages
	//
	//     Bit 4 - Enable/Disable (default) Debug Information Messages
	//
	// +kubebuilder:default:=31
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=31
	// +kubebuilder:validation:Optional
	GnssMsgNmeaUsb int `json:"gnssMsgNmeaUsb"`

	//
	// Information Message Flags for UBX Protocol on the USB (bitmask)
	// Valid range 0-31
	//
	//     Bit 0 - Enable (default) / Disable Error Information Messages
	//
	//     Bit 1 - Enable (default) / Disable Warning Information Messages
	//
	//     Bit 2 - Enable (default) / Disable Notice Information Messages
	//
	//     Bit 3 - Enable/Disable (default) Test Information Messages
	//
	//     Bit 4 - Enable/Disable (default) Debug Information Messages
	//
	// +kubebuilder:default:=31
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=31
	// +kubebuilder:validation:Optional
	GnssMsgUbxUsb int `json:"gnssMsgUbxUsb"`

	//
	// *********************** CFG-ITFM- Jamming/Interference Monitor Configuration ********************
	//

	//
	// Enable/disable interference detection
	// Valid range 0-1
	//
	//     0 - Disable interference detection
	//
	//     1 - Enable interference detection (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssIntfDetect int `json:"gnssIntfDetect"`

	//
	// Antenna setting
	// Valid range 0-2
	//
	//     0 - Unknown
	//
	//     1 - Passive
	//
	//     2 - Active
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	GnssAntSet int `json:"gnssAntSet"`

	//
	// **************************** CFG-TMODE- Survey-in Time Mode Configuration ***********************
	//

	//
	// Receiver Time Mode
	// Valid range 0-2
	//
	//     0 - Disabled
	//
	//     1 - Survey-in (default)
	//
	//     2 - Fixed Mode (true ARP position information required)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	GnssRecvTMode int `json:"gnssRecvTMode"`

	//
	// Survey-in Minimum Duration (sec)
	// Valid range
	//
	//     120 - (default)
	//
	// +kubebuilder:default:=120
	// +kubebuilder:validation:Optional
	GnssSvinMinDur int `json:"gnssSvinMinDur"`

	//
	// Survey-in Position Accuracy Limit (mm)
	// Valid range
	//
	//     100000 - (default)
	//
	// +kubebuilder:default:=100000
	// +kubebuilder:validation:Optional
	GnssSvinAccLimit int `json:"gnssSvinAccLimit"`

	//
	// **************************** GNSS Clock Out Configuration ***********************
	//

	//
	// GNSS Lock Mode
	// Valid range 0-1
	//
	//     0 - Manual
	//
	//     1 - Auto (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssLockMode int `json:"gnssLockMode"`

	//
	// GNSS Lock Threshold (nsec)
	// Valid range 0-10000
	//
	//     100 - (default)
	//
	// +kubebuilder:default:=100
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=10000
	// +kubebuilder:validation:Optional
	GnssLockTh int `json:"gnssLockTh"`

	//
	// Enable/disable the Clock Out
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	GnssClockOutEn int `json:"gnssClockOutEn"`
}

// StsConfigSpec defines the desired state of StsConfig
type StsConfigSpec struct {
	Interfaces   []StsInterfaceSpec `json:"interfaces"`
	NodeSelector map[string]string  `json:"nodeSelector,omitempty"`
	GnssSpec     StsGnssSpec        `json:"gnssSpec,omitempty"`

	// +kubebuilder:validation:Enum=T-GM.8275.1;T-BC-8275.1;T-TSC.8275.1;T-GM.8275.2;T-BC-P-8275.2;
	// +kubebuilder:default:="T-GM.8275.1"
	//Telecom G8275 Profile
	//
	// T-BC-8275.1 (default)
	//
	// T-GM.8275.1
	//
	// T-TSC.8275.1
	Mode string `json:"mode,omitempty"`

	// +kubebuilder:default:="sts-silicom"
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	//Set 1PPS Connector Mode
	//
	//1 - PPS IN
	//
	//2 - PPS OUT (default)
	//
	ModePPS int `json:"modePPS,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	ModePPS int `json:"modePPS,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:Optional
	// Set PPS OUT Source
	//
	// 1 - PLL (default)
	//
	// 2 - GPS
	//
	// 3 - IN
	SrcPPS int `json:"srcPPS,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:Optional
	// Set 10MHz Connector Mode
	//
	// 1 - 10MHz IN
	//
	// 2 - 10MHz OUT (default)
	//
	// 3 - PPS OUT
	Mode10MHz int `json:"mode10MHz,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:Optional
	// Set 10MHz OUT Source
	//
	// 1 - PLL (default)
	//
	// 2 - GPS
	//
	// 3 - IN
	Src10MHz int `json:"src10MHz,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	// +kubebuilder:validation:Optional
	// Set SyncE Recovery Clock Port
	SynceRecClkPort int `json:"synceRecClkPort,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	//Enable/disable Phy Leds Control Switch
	//
	//0 - disable Phy Leds Control Switch (default)
	//
	//1 - enable Phy Leds Control Switch
	PhyLedsCtl int `json:"phyLedsCtl,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	// Configures the synchronization network
	//
	//1 - Option 1 refers to synchronization networks designed for Europe
	//
	//2 - Option 2 refers to synchronization networks designed for United States
	SyncOption int `json:"syncOption,omitempty"`

	// +kubebuilder:default:=10
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=40
	// +kubebuilder:validation:Optional
	// Set CPU Pin for SyncE ESMC thread
	SynceCpu int `json:"synceCpu,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	//Enable/disable two-step PTP Clock
	//
	//0 - Disable two-step clock, (set one-step clock) (default)
	//
	//1 - Enable two-step clock
	TwoStep int `json:"twoStep,omitempty"`

	// +kubebuilder:default:=128
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:validation:Optional
	//Set Priority 2 for GM PTP Clock
	//
	// Valid range 0-255, smaller values indicate higher priority
	Priority2 int `json:"priority2,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	//Forwardable/Non-Forwardable Multicast Address
	//
	//0 - Non-Forwardable (default)
	//
	//1 - Forwardable
	Forwardable int `json:"forwardable,omitempty"`

	// +kubebuilder:default:=-1
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:validation:Optional
	//Trace PTP Message
	//
	//Valid range -1-10
	//
	//-1 - Disable Trace log for PTP Messages (default)
	//
	//0 - Enable Trace for all types of PTP Messages
	//
	//1 - Enable Trace for SYNC Messages
	//
	//2 - Enable Trace for DELAY_REQ Messages
	//
	//3 - Enable Trace for PEER_DELAY_REQ Messages
	//
	//4 - Enable Trace for PEER_DELAY_RESP Messages
	//
	//5 - Enable Trace for FOLLOW_UP Messages
	//
	//6 - Enable Trace for DELAY_RESP Messages
	//
	//7 - Enable Trace for PEER_DELAY_FOLLOW_UP Messages
	//
	//8 - Enable Trace for ANNOUNCE Messages
	//
	//9 - Enable Trace for SIGNAL Messages
	//
	//10 - Enable Trace for MANAGEMENT Messages
	TracePtpMsg int `json:"tracePtpMsg,omitempty"`

	// +kubebuilder:default:=23
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=36
	// +kubebuilder:validation:Optional
	//Trace module
	//
	//Valid range 0-36
	//
	//0 - Read/Write
	//
	//1 - Init
	//
	//2 - Lan
	//
	//3 - Lan Stats
	//
	//4 - Device specific interrupt
	//
	//5 - System interrupt
	//
	//6 - TS Engine interrupt
	//
	//7 - Packet interrupt
	//
	// 8 - PLL interrupt
	//
	// 9 - Signal Handler
	//
	//10 - TS Packet Stream interrupt
	//
	//11 - Transport Layer interrupt
	//
	//12 - PTP Timestamp interrupt
	//
	//13 - Packet Schedule interrupt
	//
	//14 - Main PTP Engine
	//
	//15 - PTP Best-Master-Clock related
	//
	//16 - PTP Unicast Negotiation related
	//
	//17 - PTP Unicast Discovery related
	//
	//18 - PTP Clock, Port or Stream State related
	//
	//19 - TS RECORD MGR
	//
	//20 - Socket Layer
	//
	//21 - Clock Switch
	//
	//22 - DCO MGR
	//
	//23 - Track Packet Process (default)
	//
	//24 - TOD Manager
	//
	//25 - TSIF
	//
	//26 - MSGQ
	//
	//27 - FPE
	//
	//28 - PTP Foreign Master Table
	//
	//29 - PTSF
	//
	//30 - Notify
	//
	//31 - Signal Pipe Handler
	//
	//32 - G781
	//
	//33 - PTP Timer
	//
	//34 - PTP Tlv
	//
	//35 - HO Utils
	//
	//36 - TSA
	TraceModule int `json:"traceModule,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Optional
	TraceLevel int `json:"traceLevel,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	//Configures the ESMC Mode
	//
	//1 - Manual
	//
	//2 - Auto (default)
	EsmcMode int `json:"esmcMode,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Optional
	//Configures the SSM Mode
	//
	//1 - SSM Code (default)
	//
	//2 - ESSM Code
	SsmMode int `json:"ssmMode,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Optional
	AprLevel int `json:"aprLevel,omitempty"`

	// +kubebuilder:default:=500
	// +kubebuilder:validation:Minimum=300
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:validation:Optional
	SynceHoldOff int `json:"synceHoldOff,omitempty"`

	// +kubebuilder:default:=24
	// +kubebuilder:validation:Minimum=24
	// +kubebuilder:validation:Maximum=43
	// +kubebuilder:validation:Optional
	DomainNum_8275_1 int `json:"domainNum_8275_1"`

	// +kubebuilder:default:=44
	// +kubebuilder:validation:Minimum=44
	// +kubebuilder:validation:Maximum=63
	// +kubebuilder:validation:Optional
	DomainNum_8275_2 int `json:"domainNum_8275_2"`

	// +kubebuilder:default:=4
	// +kubebuilder:validation:Minimum=4
	// +kubebuilder:validation:Maximum=23
	// +kubebuilder:validation:Optional
	DomainNum_8265_2 int `json:"domainNum_8265_2"`
}

type StsInterfaceSpec struct {
	EthName string `json:"ethName"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Optional
	// This is 1 based
	EthPort int `json:"ethPort"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	SyncE int `json:"synce,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	Ipv6 int `json:"ipv6,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	Ipv4 int `json:"ipv4,omitempty"`

	// +kubebuilder:default:=500
	// +kubebuilder:validation:Minimum=300
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:validation:Optional
	HoldOff int `json:"holdoff,omitempty"`

	// +kubebuilder:validation:Enum=Master;Slave
	// +kubebuilder:default:=Master
	// +kubebuilder:validation:Optional
	Mode string `json:"mode,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Optional
	// Configures QL enable for the interface
	// 0 - Disable QL
	// 1 - Enable QL (default)
	QlEnable int `json:"qlEnable,omitempty"`

	// +kubebuilder:default:=4
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=7
	// +kubebuilder:validation:Optional
	//Configures QL value for the interface
	//
	//if syncOption is 1 (Europe)
	//
	//===========================
	//
	//1 - QL-PRC
	//
	//2 - QL-PRTC
	//
	//3 - QL-EEC1
	//
	//4 - QL-DNU (default)
	//
	//if syncOption is 2 (United States)
	//
	//==================================
	//
	//5 - QL-PRS
	//
	//2 - QL-PRTC
	//
	//6 - QL-EEC2
	//
	//7 - QL-DUS (default)
	Ql int `json:"ql,omitempty"`
}

// StsConfigStatus defines the observed state of StsConfig
type StsConfigStatus struct {
	Nodes []string `json:"nodes,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StsConfig is the Schema for the stsconfigs API
type StsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:Required
	Spec   StsConfigSpec   `json:"spec,omitempty"`
	Status StsConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StsConfigList contains a list of StsConfig
type StsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StsConfig{}, &StsConfigList{})
}
