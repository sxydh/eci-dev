package cn.net.bhe.eciportal.service;

import cn.net.bhe.eciportal.model.dto.EciDTO;
import cn.net.bhe.eciportal.model.vo.EciVO;

import java.util.List;

public interface EciService {

    List<EciVO> queryEciList();

    void addEci(EciDTO eciDTO);

    void deleteEci(String name);

}
